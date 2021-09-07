package main

import (
	"context"
	"errors"
	"fmt"
	"keyvaluestore/store"
	"keyvaluestore/transport"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// create new in memory store
	ims := store.NewInMemoryStore()
	// create http handler
	h := transport.CreateHTTPHandler(ims)

	// launch writetofile function in a new thread
	go ims.WriteToFile()
	// create http server
	server := http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 * 1024 * 1024, // 1mb
		Handler:        h,
	}
	// create errors channel
	errs := make(chan error)

	// listen interrupt and terminate signals
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errs <- errors.New((<-done).String())
	}()

	// start http server
	go func() {
		fmt.Printf("http server listening on :8080\n")
		errs <- server.ListenAndServe()
	}()

	// wait for interrupt signals
	err := <-errs
	fmt.Printf("error: %s \n", err.Error())

	// stop http server , gracefully shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("shutting down http server failed, %s\n", err.Error())
	}

}

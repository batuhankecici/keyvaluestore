package transport

import (
	"encoding/json"
	"keyvaluestore/service"
	"net/http"
)

//Creates and returns http handler
func CreateHTTPHandler(ims service.InMemoryService) http.Handler {
	sm := http.NewServeMux()
	sm.HandleFunc("/get", createGetHandlerFunc(ims))
	sm.HandleFunc("/set", createSetHandleFunc(ims))
	sm.HandleFunc("/delete", createDeleteHandlerFunc(ims))
	sm.HandleFunc("/getall", createGetAllHandlerFunc(ims))
	return sm
}

// Get endpoint handler function
func createGetHandlerFunc(ims service.InMemoryService) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		k := req.FormValue("key")
		request := service.GetValueRequest{
			Key: k,
		}
		response := ims.GetValue(request)
		json.NewEncoder(rw).Encode(response)
	}

}

// Set endpoint handler function
func createSetHandleFunc(ims service.InMemoryService) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		k := req.FormValue("key")
		v := req.FormValue("value")
		request := service.SetValueRequest{
			Key:   k,
			Value: v,
		}
		response := ims.SetValue(request)
		json.NewEncoder(rw).Encode(response)
	}
}

// Delete endpoint handler function
func createDeleteHandlerFunc(ims service.InMemoryService) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		k := req.FormValue("key")
		request := service.GetValueRequest{
			Key: k,
		}
		response := ims.DeleteValue(request)
		json.NewEncoder(rw).Encode(response)

	}

}

// GetAll endpoint handler function
func createGetAllHandlerFunc(ims service.InMemoryService) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		response := ims.GetAll()
		json.NewEncoder(rw).Encode(response)
	}

}

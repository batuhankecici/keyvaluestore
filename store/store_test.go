package store

import (
	"bufio"
	"keyvaluestore/service"
	"os"
	"strings"
	"testing"
)

func TestSet(t *testing.T) {

	ims := NewInMemoryStore()

	k := "batuhan"
	v := "kecici"

	setValueReq := service.SetValueRequest{
		Key:   k,
		Value: v,
	}
	setValueResponse := ims.SetValue(setValueReq)

	if setValueResponse.Key != k {
		t.Error("setting error, key does not match")
	}

	if setValueResponse.Value != v {
		t.Error("setting error, value does not match")
	}
}

func TestGet(t *testing.T) {
	ims := NewInMemoryStore()
	k := "book"
	v := "kitap"
	ims.SetValue(service.SetValueRequest{
		Key:   k,
		Value: v,
	})

	getValueRequest := service.GetValueRequest{
		Key: k,
	}
	getValueResponse := ims.GetValue(getValueRequest)

	if getValueResponse.Key != k {
		t.Error("getting error, key does not match")
	}

	if getValueResponse.Value != v {
		t.Error("getting error, value does not match")
	}
}

func TestDelete(t *testing.T) {

	ims := NewInMemoryStore()
	k := "computer"
	v := "bilgisayar"
	ims.SetValue(service.SetValueRequest{
		Key:   k,
		Value: v,
	})
	getValueRequest := service.GetValueRequest{
		Key: k,
	}
	ims.DeleteValue(getValueRequest)
	getValueResponse := ims.GetValue(getValueRequest)

	if getValueResponse.Key == k {
		t.Error("delete error, still key in memory")
	}

}

func TestGetAll(t *testing.T) {
	ims := NewInMemoryStore()

	tableStruct := []struct {
		key   string
		value string
	}{
		{"ahmet", "deniz"},
		{"fatih", "terim"},
		{"key", "value"},
	}

	for _, store := range tableStruct {
		ims.SetValue(service.SetValueRequest{
			Key:   store.key,
			Value: store.value,
		})
	}

	for _, store := range tableStruct {
		getValueResponse := ims.GetValue(service.GetValueRequest{
			Key: store.key,
		})
		if getValueResponse.Key != store.key {
			t.Error("get all error, key does not find in memory")
		}
	}
}

func TestWriteToFile(t *testing.T) {
	ims := NewInMemoryStore()

	tableStruct := []struct {
		key   string
		value string
	}{
		{"ahmet", "deniz"},
		{"fatih", "terim"},
		{"key", "value"},
	}

	for _, store := range tableStruct {
		ims.SetValue(service.SetValueRequest{
			Key:   store.key,
			Value: store.value,
		})
	}

	getAllResponse := ims.GetAll()
	file, err := os.OpenFile("keyvalue-dbtest.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	for _, stores := range getAllResponse.Stores {
		for key, value := range stores {
			file.WriteString("key: " + key + " value: " + value + "\n")
		}
	}
	file.Close()

	fileRead, err := os.OpenFile("keyvalue-dbtest.txt", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer fileRead.Close()
	scanner := bufio.NewScanner(fileRead)
	for i := 0; i < 3; i++ {
		if scanner.Scan() {
			satir := scanner.Text()

			if !strings.Contains(satir, tableStruct[i].key) || !strings.Contains(satir, tableStruct[i].value) {
				t.Error("key or value not find in file")
			}
		}
	}
}

func BenchmarkSet(b *testing.B) {
	ims := NewInMemoryStore()

	k := "batuhan"
	v := "kecici"

	setValueReq := service.SetValueRequest{
		Key:   k,
		Value: v,
	}

	for i := 0; i < b.N; i++ {
		ims.SetValue(setValueReq)
	}
}

func BenchmarkGet(b *testing.B) {
	ims := NewInMemoryStore()

	k := "batuhan"
	v := "kecici"

	setValueReq := service.SetValueRequest{
		Key:   k,
		Value: v,
	}
	getValueReq := service.GetValueRequest{
		Key: k,
	}
	ims.SetValue(setValueReq)
	for i := 0; i < b.N; i++ {
		ims.GetValue(getValueReq)
	}
}

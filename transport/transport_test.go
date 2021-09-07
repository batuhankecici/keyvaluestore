package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"keyvaluestore/service"
	inMemoryStore "keyvaluestore/store"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAll(t *testing.T) {

	ims := inMemoryStore.NewInMemoryStore()
	handler := CreateHTTPHandler(ims)
	server := httptest.NewServer(handler)
	defer server.Close()

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

	k := "ball"
	v := "top"
	url := fmt.Sprintf("%s/set?key=%s&value=%s", server.URL, k, v)
	testSetHandle(t, k, v, url)
	url = fmt.Sprintf("%s/get?key=%s", server.URL, k)
	testGetHandle(t, k, v, url)
	url = fmt.Sprintf("%s/delete?key=%s", server.URL, k)
	testDeleteHandle(t, k, url)
	url = fmt.Sprintf("%s/getall", server.URL)
	testGetAllHandle(t, url)

}

func testGetHandle(t *testing.T, k, v, url string) {
	response, err := http.Get(url)
	if err != nil {
		t.Errorf("request failed %s", err.Error())
	}
	defer response.Body.Close()
	getValueResponse := &service.GetValueResponse{}
	err = json.NewDecoder(response.Body).Decode(getValueResponse)
	if err != nil {
		t.Errorf("decoding failed %s", err.Error())
	}
	if getValueResponse.Key != k {
		t.Error("key does not match")
	}
	if getValueResponse.Value != v {
		t.Error("value does not match")
	}
}

func testSetHandle(t *testing.T, k, v, url string) {
	setValueRequest := service.SetValueRequest{
		Key:   k,
		Value: v,
	}
	jsonSetValueRequest, err := json.Marshal(setValueRequest)
	if err != nil {
		t.Errorf("json marshalling error %s", err.Error())
	}
	requestBody := bytes.NewBuffer(jsonSetValueRequest)
	request, err := http.NewRequest(http.MethodPost, url, requestBody)
	if err != nil {
		t.Errorf("creating request failed: %s", err.Error())
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Errorf("Request failed: %s", err.Error())
	}

	defer response.Body.Close()

	setValueResponse := &service.SetValueResponse{}
	errs := json.NewDecoder(response.Body).Decode(setValueResponse)

	if err != nil {
		t.Errorf("decoding response failed: %s", errs.Error())
	}

	if setValueResponse.Key != k {
		t.Error("key does not match in memory")
	}

	if setValueResponse.Value != v {
		t.Error("value does not match in memory")
	}
}

func testDeleteHandle(t *testing.T, k, url string) {
	getValueRequest := service.GetValueRequest{
		Key: k,
	}
	jsonGetValueRequest, err := json.Marshal(getValueRequest)
	if err != nil {
		t.Errorf("json marshalling error: %s", err.Error())
	}
	requestBody := bytes.NewBuffer(jsonGetValueRequest)
	request, err := http.NewRequest(http.MethodPost, url, requestBody)
	if err != nil {
		t.Errorf("creating request failed: %s", err.Error())
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Errorf("request failed: %s", err.Error())
	}
	defer response.Body.Close()

	deleteValueResponse := &service.DeleteValueResponse{}
	errs := json.NewDecoder(response.Body).Decode(deleteValueResponse)

	if errs != nil {
		t.Errorf("decoding response failed: %s", errs.Error())
	}

	if !strings.Contains(deleteValueResponse.Error, "deleted") {
		t.Error("key is not in memory")
	}
}

func testGetAllHandle(t *testing.T, url string) {

	tableStruct := []struct {
		key   string
		value string
	}{
		{"ahmet", "deniz"},
		{"fatih", "terim"},
		{"key", "value"},
	}
	response, err := http.Get(url)

	if err != nil {
		t.Errorf("request failed: %s", err.Error())
	}
	defer response.Body.Close()
	getAllResponse := &service.GetAllResponse{}
	errs := json.NewDecoder(response.Body).Decode(getAllResponse)
	if errs != nil {
		t.Errorf("decoding failed: %s", errs.Error())
	}
	for _, store := range tableStruct {
		for _, value := range getAllResponse.Stores {
			//t.Error(value)
			if value[store.key] != store.value {
				t.Errorf("Get all: key does not match %s", store.key)
			}
		}
	}

}

func BenchmarkCreateHttpHandler(b *testing.B) {
	ims := inMemoryStore.NewInMemoryStore()

	for i := 0; i < b.N; i++ {
		CreateHTTPHandler(ims)
	}
}

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateCarHandler(t *testing.T) {
	jsonData := `{"id":"1","name":"toyota", "model":"corola","year":"2018","color":"black"}`
	bodyReader := strings.NewReader(jsonData)
	req, err := http.NewRequest("GET", "localhost:8080/", bodyReader)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	createCarHandler(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	// defer res.Body.Close()

	// b, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	t.Fatalf("could not read body: %v", err)
	// }

	// if string(b) != jsonData {
	// 	t.Errorf("expected %v but got %v", jsonData, string(b))
	// }
}

func TestGetCarsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/", nil)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rc := httptest.NewRecorder()

	getCarsHandler(rc, req)

	res := rc.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK!; get %v", res.Status)
	}
}

func TestGetCarHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/1", nil)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rc := httptest.NewRecorder()

	getCarHandler(rc, req)

	res := rc.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK!; get %v", res.Status)
	}
}

func TestDeleteCarHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/1", nil)

	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rc := httptest.NewRecorder()

	deleteCarHandler(rc, req)

	res := rc.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK!; get %v", res.Status)
	}
}

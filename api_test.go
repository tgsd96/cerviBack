package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mitchellh/mapstructure"

	"github.com/julienschmidt/httprouter"
	"github.com/tgsd96/cerviBack/models"
)

func TestStatusApi(t *testing.T) {
	// client := &http.Client{}
	router := httprouter.New()
	// server := httptest.NewServer(router.POST("/api/status/:image_key", StatusAPI))
	router.POST("/api/status/:image_key", StatusAPI)

	req, _ := http.NewRequest("POST", "/api/status/checkStatus.png", nil)
	var res models.Message
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Errorf("\nFailed to handle request")
	}
	payload, _ := ioutil.ReadAll(rr.Body)
	err := json.Unmarshal(payload, &res)
	if err != nil {
		t.Errorf("\nError: %s", err.Error())
	}
	if res.Action == "image_not_found" {
		t.Fatalf("\n Image not found")
	} else {
		t.Logf("\n Action Received : ", res.Action)
	}
	var result models.Result
	err = mapstructure.Decode(res.Data, &result)
	if err != nil {
		t.Errorf("\n Unable to decode to results: %s", err.Error())
	}
	if result.Status == "" {
		t.Errorf("\n Received no status")
	} else {
		t.Logf("\n Received status: %s, type1: %f, type2: %f, type3:%f", result.Status, result.Type1, result.Type2, result.Type3)

	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
Ref:
https://play.golang.org/p/eQv_QouLnME
*/

func Test(t *testing.T) {

	s := `{"request":{"uid": "1","object":{"metadata":{"labels":{"billing":"y"}}}}}`

	bodyReader := strings.NewReader(s)
	req := httptest.NewRequest("POST", "/Validate", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Validate(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	expected := `{"apiVersion": "admission.k8s.io/v1","kind": "AdmissionReview","response": {"allowed": true, "uid": 1, "status": {"message": "valid"}}})`

	if string(body) != expected {
		t.Fatalf("\nexpected: ->%s\ngot: %s\n", expected, body)
	}

	fmt.Printf("%v\n", string(body))
}
func TestBroken(t *testing.T) {

	s := `{"request":{"uid": "1","object":{"metadata":{"labels":{"stuff":"y"}}}}}`

	bodyReader := strings.NewReader(s)
	req := httptest.NewRequest("POST", "/Validate", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	Validate(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	expected := `{"apiVersion": "admission.k8s.io/v1","kind": "AdmissionReview","response": {"allowed": false, "uid": 1, "status": {"message": "Not valid"}}})`
	if string(body) != expected {
		t.Fatalf("\nexpected: ->%s\ngot: %s\n", expected, body)
	}

	fmt.Printf("%v\n", string(body))

}

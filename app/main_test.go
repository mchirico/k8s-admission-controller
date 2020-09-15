package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
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
	req := httptest.NewRequest("POST", "/validate", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	validate(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	expected := `{"apiVersion": "admission.k8s.io/v1","kind": "AdmissionReview","response": {"allowed": true, "uid": 1, "status": {"message": "valid"}}})`

	if !reflect.DeepEqual(string(body), expected) {
		t.Fatalf("\nexpected: ->%s\ngot: %s\n", expected, body)
	}

	fmt.Printf("%v\n", string(body))
}
func TestBroken(t *testing.T) {

	s := `{"request":{"uid": "1","object":{"metadata":{"labels":{"stuff":"y"}}}}}`

	bodyReader := strings.NewReader(s)
	req := httptest.NewRequest("POST", "/validate", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	validate(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	expected := `{"apiVersion": "admission.k8s.io/v1","kind": "AdmissionReview","response": {"allowed": false, "uid": 1, "status": {"message": "Not valid"}}})`
	if !reflect.DeepEqual(string(body), expected) {
		t.Fatalf("\nexpected: ->%s\ngot: %s\n", expected, body)
	}

	fmt.Printf("%v\n", string(body))

}

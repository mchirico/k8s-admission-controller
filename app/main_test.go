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
	req := httptest.NewRequest("POST", "/validate", bodyReader)
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()
	validate(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v\n",string(body))
}

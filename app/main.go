package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	//"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type AutoGenerated struct {
	Request struct {
		UID    string `json:"uid"`
		Object struct {
			Metadata struct {
				Labels struct {
					Billing string `json:"billing"`
				} `json:"labels"`
			} `json:"metadata"`
		} `json:"object"`
	} `json:"request"`
}

func checkJSON(r *http.Request) (AutoGenerated, error) {



		var m AutoGenerated
		err := json.NewDecoder(r.Body).Decode(&m)

		return m,err
		//fmt.Printf("%s: %s\n", m.Request.UID, m.Request.Object.Metadata.Labels.Billing)

}

func checkHeader(r *http.Request) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return errors.New(msg)
		}
		return nil
	}
	return errors.New("Header blank")
}



func validate(w http.ResponseWriter, r *http.Request) {

	err := checkHeader(r)
	if err != nil {
		return
	}

	dump, _ := httputil.DumpRequest(r, true)
	data,err := checkJSON(r)
	if err != nil {
		return
	}
	msg := fmt.Sprintf("\nhere:\n%v:",string(dump))
	log.Println("-->",data.Request.UID)
	io.WriteString(w, msg)
}

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/validate", validate)

	log.Printf("About to listen on 5000. Go to https://%s:5000/", name)
	err = http.ListenAndServeTLS(":5000", "./certs/server_certificate.pem", "./certs/server_key.pem", nil)
	log.Fatal(err)
}

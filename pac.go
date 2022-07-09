package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

var pacContent []byte

func pacServer(fileName string, port int) {
	var err error
	pacContent, err = ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	p := path.Base(fileName)
	log.Printf("pac http server: http://localhost:%d/%s", port, p)

	http.HandleFunc("/"+p, pacHandler)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pacHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write(pacContent)
	if err != nil {
		log.Println(err)
	}
}

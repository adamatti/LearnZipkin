package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	// "github.com/thedevsaddam/gojsonq"
	"github.com/Jeffail/gabs"
)

func changeJson(body []byte) {
	gabs.ParseJSON(body)
}

func requestCallback(req *http.Request) {
	var buf bytes.Buffer
	tee := io.TeeReader(req.Body, &buf)
	defer req.Body.Close()

	body, _ := ioutil.ReadAll(tee)
	log.Println("RequestCallback", string(body))

	req.Body = ioutil.NopCloser(
		bytes.NewReader(
			buf.Bytes(),
		),
	)
}

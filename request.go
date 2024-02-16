package main

import (
	"io"
	"net/http"
)

func post_request(url string, headers map[string][]string, body io.Reader) (io.ReadCloser, int) {
	r, err := http.NewRequest("POST", url, body)

	if err != nil {
		panic(err)
	}

	r.Header = headers

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	return res.Body, res.StatusCode
}

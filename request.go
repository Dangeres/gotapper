package main

import (
	"io"
	"net/http"
)

func post_request(url string, headers map[string][]string, body io.Reader) ([]byte, int) {
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

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return bodyBytes, res.StatusCode
}

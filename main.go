package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func main() {
	const url = "https://hard.apitapper.ru/api/order/get"
	const shop_token = "stiks-4114"

	for i := 1; i < 40; i++ {
		fmt.Println(i)

		browserid := randSeq(15)
		fmt.Println(browserid)

		body_raw := RequestTapper{
			TableId:   strconv.Itoa(i),
			Domen:     shop_token,
			Guest:     rand.Int(),
			BrowserId: BrowserId{BrowserId: browserid, Plugin: "fingerprint"},
		}

		body, _ := json.Marshal(body_raw)

		result, status_code := post_request(
			url,
			map[string][]string{
				"Content-Type": {"application/json, text/plain, */*"},
				"User-Agent":   {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36"},
			},
			bytes.NewBuffer(body),
		)

		defer result.Close()

		fmt.Println(status_code)

		response := &ResponseTapper{}
		derr := json.NewDecoder(result).Decode(response)

		if derr != nil {
			panic(derr)
		}

		if len(response.Result.OrdersElementAll) > 0 {
			fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<")
		}

		fmt.Println(response)

		if len(response.Result.OrdersElementAll) > 0 {
			fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<")
		}

		fmt.Println("==========================")
	}

	fmt.Println("okay")
}

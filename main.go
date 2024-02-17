package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const folderOrder = "order/"
const folderPosition = "position/"
const fileSettings = "settings.json"

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

func changePositions(data ResponseTapper) {
	for _, p := range data.Result.OrdersElementAll {
		filePath := strings.Join([]string{folderPosition, p.PositionId, ".json"}, "")

		dataFile := readJson(filePath)

		position := CustomPositions{}

		json.Unmarshal(dataFile, &position)

		if position.Orders == nil {
			position.Orders = make(map[string]CustomPriceAmount)
		}

		position.Orders[p.OrderId] = CustomPriceAmount{Price: p.Price, Amount: p.Amount}
		position.PositionId = p.PositionId
		position.Title = p.Name

		b, _ := json.Marshal(position)

		writeJson(filePath, []byte(b))
	}
}

func main() {
	const url = "https://hard.apitapper.ru/api/order/get"

	var settings Settings
	json.Unmarshal(readJson(fileSettings), &settings)

	if settings.IsProcessData {
		entries, _ := os.ReadDir(folderOrder)

		for _, f := range entries {
			fmt.Printf("Обрабатываем позиции файла заказа %s", f.Name())

			dataFile := readJson(strings.Join([]string{folderOrder, f.Name()}, ""))

			response := ResponseTapper{}

			derr := json.Unmarshal(dataFile, &response)

			if derr != nil {
				panic(derr)
			}

			changePositions(response)
		}
	}

	for i := 1; i < 60; i++ {
		fmt.Println(i)

		browserid := randSeq(15)
		fmt.Println(browserid)

		body_raw := RequestTapper{
			TableId:   strconv.Itoa(i),
			Domen:     settings.ShopToken,
			Guest:     rand.Int(),
			BrowserId: RequestBrowserId{BrowserId: browserid, Plugin: "fingerprint"},
		}

		body, _ := json.Marshal(body_raw)

		body_result, status_code := post_request(
			url,
			map[string][]string{
				"Content-Type": {"application/json, text/plain, */*"},
				"User-Agent":   {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36"},
			},
			bytes.NewBuffer(body),
		)

		fmt.Println(status_code)

		response := ResponseTapper{}

		derr := json.Unmarshal(body_result, &response)

		if derr != nil {
			panic(derr)
		}

		if len(response.Result.OrdersElementAll) > 0 {
			fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<")

			fmt.Println(string(body_result))

			writeJson(strings.Join([]string{folderOrder, response.Result.Orders[0].OrderId, ".json"}, ""), body_result)

			if settings.IsProcessData {
				changePositions(response)
			}

			fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<")
		} else {
			fmt.Printf("На столике %d никто не сидит\n", i)
		}

		fmt.Println("==========================")
	}

	fmt.Println("program has been finished")
}

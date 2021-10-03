package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hsmtkk/go-call-alpha-vantage-api/parse"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: program symbol")
	}
	symbol := os.Args[1]

	apiKey, err := readAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	json, err := getStock(symbol, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	parse.ParseJSON(json)
}

func getStock(symbol, apiKey string) ([]byte, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s", symbol, apiKey)

	// HTTPクライアント
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスを全て読み出す
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

	// resp.Body.Close() はココで実行される
}

func readAPIKey() (string, error) {
	apiKey := os.Getenv("API_KEY")
	// API_KEY環境変数が定義されてなかったらエラーを返す
	if apiKey == "" {
		return "", fmt.Errorf("API_KEY environment variable is not defined")
	}
	return apiKey, nil
}

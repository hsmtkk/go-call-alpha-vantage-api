package parse_test

import (
	"log"
	"os"
	"testing"

	"github.com/hsmtkk/go-call-alpha-vantage-api/parse"
)

func TestParse(t *testing.T) {
	json, err := os.ReadFile("./ibm-stock.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := parse.ParseJSON(json); err != nil {
		log.Fatal(err)
	}
}

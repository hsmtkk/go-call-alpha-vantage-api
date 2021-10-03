package parse

import (
	"encoding/json"
	"fmt"
)

type Schema struct {
	MetaData   map[string]string     `json:"Meta Data"`
	TimeSeries map[string]StockValue `json:"Time Series (5min)"`
}

type StockValue map[string]string

func ParseJSON(js []byte) error {
	var sch Schema
	if err := json.Unmarshal(js, &sch); err != nil {
		return err
	}
	fmt.Println(sch.MetaData["1. Information"])
	for k, v := range sch.TimeSeries {
		fmt.Println(k, v)
		break
	}
	return nil
}

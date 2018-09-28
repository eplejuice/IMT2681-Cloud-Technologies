package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiRoot = "https://api.exchangeratesapi.io/latest"

type exRates struct {
	Rates []map[string]string `json:"rates"`
}

func exRate() []map[string]string {
	resp, err := http.Get(apiRoot)
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	exR := &exRates{}

	err = json.NewDecoder(resp.Body).Decode(exR)
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}

	return exR.Rates
}

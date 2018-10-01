package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type exR struct {
	Rates map[string]float64 `json:"rates"`
}

const apiRoot = "https://api.exchangeratesapi.io/latest"

func exRate() map[string]float64 {
	resp, err := http.Get(apiRoot)
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	var e exR

	tmp, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(tmp, &e)
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}

	return e.Rates
}

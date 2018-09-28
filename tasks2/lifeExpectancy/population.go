package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiRoot = "http://api.population.io:80/1.0/life-expectancy/total/"

//LifeExp cointains a response for life expectancy endpoint
type LifeExp struct {
	Dob     string
	Country string
	TotalLE float32 `json:"total_life_expectancy"`
	Sex     string
}

func getLifeEx(s, c, d string) float32 {

	sex := s
	country := c
	dob := d

	url := apiRoot + sex + "/" + country + "/" + dob + "/"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: problem", err)
		return 0
	}
	defer resp.Body.Close()

	le := &LifeExp{}

	err = json.NewDecoder(resp.Body).Decode(le)
	if err != nil {
		fmt.Println("Error: problem", err)
		return 0
	}

	return le.TotalLE
}

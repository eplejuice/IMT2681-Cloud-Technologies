package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiRoot = "https://api.github.com/repos/NZDIS/CameraView/languages"

type returnObj struct {
	Cplusplus float32 `json:"C++"`
	Java      float32 `json:"Java"`
	Shell     float32 `json:"Shell"`
}

func githubapi() float32 {
	resp, err := http.Get(apiRoot)
	if err != nil {
		fmt.Println("Error", err)
		return 0
	}

	rObj := &returnObj{}

	err = json.NewDecoder(resp.Body).Decode(rObj)
	if err != nil {
		fmt.Println("Error", err)
		return 0
	}

	return rObj.Cplusplus
}

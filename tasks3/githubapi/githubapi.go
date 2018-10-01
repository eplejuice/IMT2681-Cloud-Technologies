package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiRoot = "https://api.github.com/repos/"

func githubapi(user string, repo string) {

	fmt.Println(user)
	fmt.Println(repo)

	language := "C++"
	api := apiRoot + user + "/" + repo + "/" + "languages"

	fmt.Println(api)

	resp, err := http.Get(api)
	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	var rObj map[string]float32

	tmp, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(tmp, &rObj)

	if err != nil {
		fmt.Println("Error", err)
		panic(err)
	}

	_, tmpp := rObj[language]
	if !tmpp {
		fmt.Println("This project is not written in ", language)
		return
	}

	for k, v := range rObj {
		if k == language {
			fmt.Println(v)
		}
	}
}

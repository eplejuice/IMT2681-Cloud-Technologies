package main

import "net/http"

func main() {
	// Task 6
	http.HandleFunc("/hello/", webserver)
	http.HandleFunc("/", nonHello)
	http.HandleFunc("/hello", nonName)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

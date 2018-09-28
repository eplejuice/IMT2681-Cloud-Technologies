package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func webserver(w http.ResponseWriter, r *http.Request) {
	message := strings.SplitAfter(r.URL.Path, "/")
	msg := message[2]
	if isNumber(msg) {
		badRequest(w)
	} else {
		w.Write([]byte("Hello " + msg))
	}
}

func nonHello(w http.ResponseWriter, r *http.Request) {
	status := http.StatusBadRequest
	http.Error(w, http.StatusText(status), status)
}

func nonName(w http.ResponseWriter, r *http.Request) {
	status := http.StatusBadRequest
	http.Error(w, http.StatusText(status), status)
}

func isNumber(t string) bool {
	_, err := strconv.Atoi(t)
	if err == nil {
		return true
	}
	return false
}

func badRequest(w http.ResponseWriter) {
	status := http.StatusBadRequest
	http.Error(w, http.StatusText(status), status)
	fmt.Fprint(w, "Dont input numbers")
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Book struct {
	Id     string
	Year   string
	Title  string
	Author string
}

var bookarray = []Book{

	Book{
		Id:     "1",
		Year:   "1",
		Title:  "bible",
		Author: "jesus",
	},
	Book{
		Id:     "2",
		Year:   "450",
		Title:  "ArtofWar",
		Author: "SunTzu",
	},
}

func main() {
	http.HandleFunc("/books/", getById)
	http.HandleFunc("/books/3", biggerThanYouThough)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

var nummerPath, _ = regexp.Compile("[3-9]")

func getById(w http.ResponseWriter, r *http.Request) {

	/*
		fmt.Println(r.URL.Path)
		switch {
		case nummerPath.MatchString(r.URL.Path):
			hdlNmmPath(w, r)
		}
	*/

	temp := strings.SplitAfter(r.URL.Path, "/")
	rId := temp[2]
	if !isNumber(rId) {
		badRequest(w)
	} else {
		temp, err := strconv.Atoi(rId)
		if err != nil {
			panic(err)
		}
		if !isWithinRange(temp) {
			indexNotFound(w)
		} else {
			rIddd := getIdByIndex(rId)
			r.Header.Set("Content-Type", "application/json")
			bookJson, err := json.Marshal(bookarray[rIddd])
			if err != nil {
				fmt.Fprintf(w, "Error", err)
			}

			w.Write(bookJson)
		}
	}
}

func getIdByIndex(j string) int {
	fmt.Println("j = ", j)
	for i := 0; i <= len(bookarray); i++ {
		fmt.Println("i = ", i)
		fmt.Println("Id = ", bookarray[i].Id)
		if bookarray[i].Id == j {
			return i
		}
	}
	return 0
}

func isNumber(t string) bool {

	_, err := strconv.Atoi(t)
	if err == nil {
		return true
	}
	return false
}

func isWithinRange(t int) bool {

	if t > 0 && t < 3 {
		return true
	}
	return false
}

func badRequest(w http.ResponseWriter) {
	status := http.StatusBadRequest
	http.Error(w, http.StatusText(status), status)
}

func biggerThanYouThough(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books/3" {
		badRequest(w)
		return
	}
	fmt.Fprint(w, "DONT TYPE 3")
}

func indexNotFound(w http.ResponseWriter) {
	badRequest(w)
	fmt.Fprint(w, "Book with that ID is not present")
}

/*
func hdlNmmPath(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on numeric path handler..")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}
*/

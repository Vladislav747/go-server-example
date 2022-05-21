package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/", homepage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func homepage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage")
}

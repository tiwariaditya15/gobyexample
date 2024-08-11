package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)

var url string = "https://jsonplaceholder.typicode.com/todos/"


func greeter(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err1 := ioutil.ReadAll(resp.Body)

	if err1 != nil {
		log.Fatal("Failed to read json response", err1)
	}

	fmt.Fprintf(w, string(respBody))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			if name == "Accept-Language" {
				fmt.Fprintf(w, "%v is %v\n", name, "aditya")
			} else {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	}
}


func main() {
	http.HandleFunc("/healthcheck", greeter)
	http.HandleFunc("/headersinfo", headers)

	http.ListenAndServe(":8090", nil)
}
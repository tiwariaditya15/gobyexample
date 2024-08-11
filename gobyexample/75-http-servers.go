package main

import (
	"net/http"
	"fmt"
)


func greeter(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello there!")
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
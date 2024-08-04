package main

import (
	"net/http"
	// "bufio"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
)

var url string = "https://jsonplaceholder.typicode.com/todos/"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response Http Status>>>", resp.Status)


	// scanner := bufio.NewScanner(resp.Body)

	// for i := 0; scanner.Scan() && i < 5; i++ {
	// 	fmt.Println(scanner.Text())
	// }


	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }


	// reading response body
	respBody, err1 := ioutil.ReadAll(resp.Body)

	if err1 != nil {
		log.Fatal("Failed to read json response", err1)
	}

	// var result map[string]interface{} // for objects
	var result []interface{} // for arrays 
	// note: using interface{} to avoid the typing the response else should use struct types
	if err2 := json.Unmarshal(respBody, &result); err2 != nil {
		log.Fatal("Failed to unmarshal:", err2)
	}

	fmt.Println("Result>>>", result)
}
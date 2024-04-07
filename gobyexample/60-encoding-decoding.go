package main

import (
	"fmt"
	b64 "encoding/base64"
)

func main() {
	// s := "abc123!?$*&()'-=@~"
	// es := b64.StdEncoding.EncodeToString([]byte(s))
	// fmt.Println(es)

	url := "?name=aditya tiwari&city=palghar mumbai"
	ues := b64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(ues)


	uds, _ := b64.URLEncoding.DecodeString(ues)
	fmt.Println(string(uds))
}
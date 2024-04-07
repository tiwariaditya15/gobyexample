package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://deli:128812@host.com:5432/path?page=0&limit=10&channel=web,app,infra#fa"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	password, _ := u.User.Password()
	fmt.Println(password)


	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host>>>", host, "port>>>", port)


	fmt.Println(u.Path)
	fmt.Println(u.Fragment)


	fmt.Println(u.RawQuery)

	queryParams, _ := url.ParseQuery(u.RawQuery)

	fmt.Println(queryParams)
	for a, b := range queryParams["channel"] {
		fmt.Println("a>>>", a, "b>>>", b)
	}
}
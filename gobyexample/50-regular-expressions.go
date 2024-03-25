package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	r, _ := regexp.Compile("a([a-z]+)ya")
	fmt.Println(r.MatchString("aditya"))
	fmt.Println(r.FindString("tiwarin zaditydaa"))
	fmt.Println(r.FindStringIndex("tiwarin zaditya"))
	fmt.Println(r.FindStringSubmatch("tiwarin zaditya"))
	fmt.Println(r.FindStringSubmatchIndex("tiwarin zaditya"))
	fmt.Println(r.FindAllString("tiwarin zaditya sdadityasd aditya", 2)) // second param is number of macthes you want
	fmt.Println(r.FindAllStringSubmatchIndex("tiwarin zaditya sdadityasd", -1))
	fmt.Println(r.Match([]byte("tiws aditya")))
	fmt.Println(r.ReplaceAllString("an aditya", "some rando name"))

	in := []byte("an aditya")
	r = regexp.MustCompile("a([a-z]+)ya")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(out)
}
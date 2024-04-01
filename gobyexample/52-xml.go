package main

import (
	"fmt"
	"encoding/xml"
)

type Plant struct {
	XMLName	xml.Name `xml:"plant"`
	Id		int		 `xml:"id,attr"`	
	Name 	string	 `xml:"name"`
	Origin	[]string `xml:"origin"`
}

func main() {
	coffee := &Plant{Id:23,Name:"Coffee"}
	coffee.Origin = []string{"Ethiopia","Brazil"}

	out, _ := xml.MarshalIndent(coffee, "", " ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	fmt.Println("================Unmarshalling================")
	
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	
	fmt.Println("================Nesting================")

	tomato := &Plant{Id:81,Name:"Tomato"}
	tomato.Origin = []string{"Cali", "Mexico"}

	type Nesting struct {
		XMLName	xml.Name `xml:"trees"`
		Plants	[]*Plant `xml:"parent>child>aditya>plant"` // gives you hierarchy that you want 
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	nestedXml, _ := xml.MarshalIndent(nesting, " ", " ")
	fmt.Println(xml.Header, string(nestedXml))
}
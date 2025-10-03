package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN string `xml:"isbn,attr"`
	Title string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Pseudo string `xml:"pseudo"`
	PseudoAttr string `xml:"pseudoattr,attr"`
}

func main() {
	book := Book{
		ISBN: "584-75-437-587-34-32",
		Title: "Go Bootcamp",
		Author: "Ashish",
		Pseudo: "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		log.Fatalln("Error marshalling XML: ", err)
		return
	}
	fmt.Println(string(xmlDataAttr))
}
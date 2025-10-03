# 96 - XML
XML, which stands for Extensible Markup Language, is a Markup Language used for encoding documents in a format that is both human readable and machine readable.

It is widely used for data interchange between systems and for configuration files.

Go provides the encoding XML package to handle XML data. This package offers functions to encode and deconde XML data.

Before JSON became popular with Rest APIs, it was XML that was the first choice for transmitting data.

To convert go data structures into XML, we use the XML.Marshall function and to convert XML data to into go data structure, we use XML.Unmarshall().

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age"`
	City string `xml:"city"`
	Email string `xml:"email"`
}

func main() {
	person := Person{Name: "John", Age: 34, City: "London", Email: "somelahmed.prog@gmail.com"}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling XML: ", err)
		return
	}
	fmt.Println(string(xmlData))
}
```
```bash
<person><name>John</name><age>34</age><city>London</city><email>somelahmed.prog@gmail.com</email></person>
```

----------------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age"`
	City string `xml:"city"`
	Email string `xml:"email"`
}

func main() {
	person := Person{Name: "John", Age: 34, City: "London", Email: "somelahmed.prog@gmail.com"}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling XML: ", err)
		return
	}

	fmt.Println("XML Data: ", string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		fmt.Println("Error marshalling XML: ", err)
		return
	}
	fmt.Println("XML data with indent: ", string(xmlData1))
}
```
```bash
XML Data:  <person><name>John</name><age>34</age><city>London</city><email>somelahmed.prog@gmail.com</email></person>
XML data with indent:  <person>
 <name>John</name>
 <age>34</age>
 <city>London</city>
 <email>somelahmed.prog@gmail.com</email>
</person>
```

**----------------------------------------------------------------------------------------------------------------------------------**

## Unmarshalling XML data
```go
package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age"`
	City string `xml:"city"`
	Email string `xml:"email"`
}

func main() {
	XMLRawData := `<person><name>Jane</name><age>25</age></person>`
	var personXML Person

	err := xml.Unmarshal([]byte(XMLRawData), &personXML)
	if err != nil {
		log.Fatalln("Error Unmarshalling XML: ", err)
		return
	}
	fmt.Println(personXML)
	fmt.Println(personXML.XMLName)
}
```
```bash
{{ person} Jane 25  }
{ person}
```

**----------------------------------------------------------------------------------------------------------------------------------**

## using omitempty
```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age"`
	City string `xml:"city"`
	Email string `xml:"email"`
}

func main() {
	// person := Person{Name: "John", Age: 34, City: "London", Email: "somelahmed.prog@gmail.com"}
	person := Person{Name: "John", Email: "somelahmed.prog@gmail.com"}

	xmlData, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		fmt.Println("Error Marshalling XML: ", err)
		return
	}
	fmt.Println("xml Data: ", string(xmlData))


}
```
```bash
xml Data:  <person>
 <name>John</name>
 <age>0</age>
 <city></city>
 <email>somelahmed.prog@gmail.com</email>
</person>
```

----------------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age,omitempty"`
	City string `xml:"city,omitempty"`
	Email string `xml:"email"`
}

func main() {
	// person := Person{Name: "John", Age: 34, City: "London", Email: "somelahmed.prog@gmail.com"}
	person := Person{Name: "John", Email: "somelahmed.prog@gmail.com"}

	xmlData, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		fmt.Println("Error Marshalling XML: ", err)
		return
	}
	fmt.Println("xml Data: ", string(xmlData))


}
```
```bash
xml Data:  <person>
 <name>John</name>
 <email>somelahmed.prog@gmail.com</email>
</person>
```

----------------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age"`
	City string `xml:"city"`
	Email string `xml:"-"`
}

func main() {
	person := Person{Name: "John", Age: 34, City: "London", Email: "somelahmed.prog@gmail.com"}

	xmlData, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		fmt.Println("Error Marshalling XML: ", err)
		return
	}
	fmt.Println("xml Data: ", string(xmlData))


}
```
```bash
xml Data:  <person>
 <name>John</name>
 <age>34</age>
 <city>London</city>
</person>
```

**----------------------------------------------------------------------------------------------------------------------------------**

## Handling nested XML - Marshalling
```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Address struct {
	City string `xml:"city"`
	State string `xml:"State"`
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age"`
	Address Address `xml:"address"`
	Email string `xml:"-"`
}

func main() {
	person := Person{Name: "John", Age: 34, Address: Address{City: "London", State: "CA"}, Email: "somelahmed.prog@gmail.com"}

	xmlData, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		fmt.Println("Error Marshalling XML: ", err)
		return
	}
	fmt.Println("xml Data: ", string(xmlData))


}
```
```bash
xml Data:  <person>
 <name>John</name>
 <age>34</age>
 <address>
  <city>London</city>
  <State>CA</State>
 </address>
</person>
```

**----------------------------------------------------------------------------------------------------------------------------------**

## Handling nested XML - Unmarshalling
```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Address struct {
	City string `xml:"city"`
	State string `xml:"state"`
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age"`
	Address Address `xml:"address"`
	Email string `xml:"-"`
}

func main() {
	xmlRawData := `<person><name>John</name><age>25</age><address><city>Newyork</city><state>CA</state></address></person>`

	var personXML Person
	err := xml.Unmarshal([]byte(xmlRawData), &personXML)
	if err != nil {
		fmt.Println("Error Unmarshalling XML: ", err)
		return
	}
	fmt.Println("Decode/Unmarshall Data: ", personXML)
}
```
```bash
Decode/Unmarshall Data:  {{ person} John 25 {Newyork CA} }
```

**----------------------------------------------------------------------------------------------------------------------------------**

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Address struct {
	City string `xml:"city"`
	State string `xml:"state"`
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name string `xml:"name"`
	Age int `xml:"age"`
	Address Address `xml:"address"`
	Email string `xml:"-"`
}

func main() {
	xmlRawData := `<person><name>John</name><age>25</age><address><city>Newyork</city><state>CA</state></address></person>`

	var personXML Person
	err := xml.Unmarshal([]byte(xmlRawData), &personXML)
	if err != nil {
		fmt.Println("Error Unmarshalling XML: ", err)
		return
	}
	fmt.Println("Decode/Unmarshall Data: ", personXML)
	fmt.Println("Local: ", personXML.XMLName.Local)
	fmt.Println("Space: ", personXML.XMLName.Space)
}
```
```bash
Decode/Unmarshall Data:  {{ person} John 25 {Newyork CA} }
Local:  person
Space:
```

**----------------------------------------------------------------------------------------------------------------------------------**

## attribute -> <book isbn="dsfsjflksf" color="blue">
```go
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
}

func main() {
	book := Book{
		ISBN: "584-75-437-587-34-32",
		Title: "Go Bootcamp",
		Author: "Ashish",
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		log.Fatalln("Error marshalling XML: ", err)
		return
	}
	fmt.Println(string(xmlDataAttr))
}
```
```bash
<book isbn="584-75-437-587-34-32" title="Go Bootcamp" author="Ashish"></book>
```

**----------------------------------------------------------------------------------------------------------------------------------**

## adding attributes value to child elements - need to create separate struct for those child elements if we do like below than all attributed value will be added with XMLName.
```go
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
```
```bash
<book isbn="584-75-437-587-34-32" title="Go Bootcamp" author="Ashish" pseudoattr="Pseudo Attribute">
 <pseudo>Pseudo</pseudo>
</book>
```

**----------------------------------------------------------------------------------------------------------------------------------**

## why we are concern about XML?
why should we concerned about XML, when most of the work is done using JSON?
and we are moving forward to towards gRPC as well?

**Rest API use JSON snd gRPC to transfer data in a binary format**

XML is widely used in real world scenarios, especially in industries where data interchange and configuration management are critical.

- Native android development uses XML intensively.
- Web services and APIs use XML
- Before Rest APIs were popular, we used to have Soap. Soap was another standard of making APIs. It actually stands for Simple Object Access Protocol, and soap is a Protocol for exchanging structured information in a web services using XML for message format. In JSON, we are using JSON Object and in XML, we would use message. And even till today, many enterprise systems ans legacy applications still use Soap based web services. So people who have knowledge about XML, are the people who get chosen to work on those legacy systems and legacy services.
- Some examples that use Soap and XML are payment gateways, financial services, and telecom APIs
- Now, when we come to Rest APIs, although JSON is more popular in modern APIs, XML is still used in many APIs for data interchange, especially in older systems when dealing with complex data structures
- Another usage of XML that you will see is in Spring Framework.
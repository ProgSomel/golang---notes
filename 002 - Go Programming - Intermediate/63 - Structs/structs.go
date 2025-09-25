package main

import "fmt"

type Person struct {
		firstName string
		lastName string
		age int
		address Address
		PhoneHomeCell
}

type Address struct{
	city string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}



func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main(){
	
	p := Person {
		firstName: "John",
		lastName: "Doe",
		age: 30,
		address: Address{
			city: "Dhaka",
			country: "Bangladesh",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "46576",
			cell: "4556778",
		},
	}

	p2 := Person{
		firstName: "John",
		lastName: "Doe",
		age: 30,
		PhoneHomeCell: PhoneHomeCell{
			home: "46576",
			cell: "4556778",
		},
	}
	p2.address.city = "Dhaka"
	p2.address.country = "Bangladesh"
	
	fmt.Println(p==p2)
}
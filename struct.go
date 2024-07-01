package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	age       int // by default is zero
}

type Contact struct {
	email   string
	zipCode int
}
type Address struct {
	street string
	city   string
}

type Employee struct {
	Person_details  person
	Contact_details Contact
	Address_details Address
}

func main() {
	var p1 person
	fmt.Println(p1)
	p1.FirstName = "divyanshu"
	p1.LastName = "kumar"
	p1.age = 25
	fmt.Println(p1)

	person1 := person{FirstName: "divyanshu", LastName: "old", age: 55}
	fmt.Println(person1)

	// new keyword
	p2 := new(person)
	p2.FirstName = "divyanshu"
	p2.LastName = "kumar"
	p2.age = 25

	// creating employee

	employee1 := Employee{
		Person_details: person{
			FirstName: "divyanshu",
			LastName:  "kumar",
			age:       25,
		},
		Contact_details: Contact{
			email:   "123@gmail.com",
			zipCode: 123,
		},
		Address_details: Address{
			street: "123 street",
			city:   "New York",
		},
	}

	// can also be accesses using dot operator

	employee1.Person_details.FirstName = "divyanshu new"
	fmt.Println(employee1)
	fmt.Println(employee1.Person_details.FirstName)

}

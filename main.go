package main

import "fmt"

type contact struct {
	phone string
	email string
}
type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	str := "Hello"

	aziz := person{
		firstName: "Aziz",
		lastName:  "Rezgui",
		contact: contact{
			phone: "58222897",
			email: "azizrezgui4@gmail.com",
		},
	}

	aziz.updateName("Zizou")

	editStr(&str)
	fmt.Println(str)
}

func editStr(str *string) {
	*str += "asba"
}
func (p *person) updateName(newFirstName string) {

	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

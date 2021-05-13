package main

import "fmt"

// Data structure in GO to store related data, kind of like dict in python
type person struct {
	firstname   string
	lastname    string
	contactinfo // embed another struct
}

type contactinfo struct {
	email string
	zip   int
}

func main() {
	//embedStruct()
	PointerExplaination()

}

// function to declare diffrent struct type
func declareTypeStruct() {
	// Declaration 1, you can also do without telling keyname and it will auto assign based on order
	pname := person{firstname: "Ishan", lastname: "Kumar"}
	fmt.Println(pname)

	// Declaration 2
	var pname2 person //Auto assign empty stuct by default
	pname2.firstname = "Ashish"
	pname2.lastname = "Nigam"
	fmt.Printf("%v", pname2)
}

// function to embed struct within each other
func embedStruct() {
	embStruct := person{
		firstname: "peter",
		lastname:  "pan",
		contactinfo: contactinfo{ // assign values to embedded struct
			email: "peter@test.com",
			zip:   12345,
		},
	}
	fmt.Printf("%+v", embStruct)
}

// print a struct of type person
func (p person) print() {
	fmt.Printf("Struct to print --> %+v", p)
}

// function to Update a value using pointer
// *Person can take in both the value or the pointer
func (pounterToPerson *person) updateName(newlasttname string) {
	(*pounterToPerson).lastname = newlasttname // Update the lastname using the actual location
}

// PointerExplaination
func PointerExplaination() {
	embStruct := person{
		firstname: "peter",
		lastname:  "pan",
		contactinfo: contactinfo{ // assign values to embedded struct
			email: "peter@test.com",
			zip:   12345,
		},
	}

	// this will update the copy of embStruct and not the actual value, pointers concept
	// embStruct.updateName("parker") // updated the function to make it work
	// embStruct.print()

	// to actually update struct value, use struct location
	// then update the value using that location
	embStructPointer := &embStruct
	//embStructPointer.updateName("parker")

	// You cal also call it durectly with var name
	fmt.Println("Location of a variable embtruct -->", embStructPointer)
	fmt.Println("Value of a variable embtruct -->", *&embStructPointer)

	embStruct.updateName("Quil")
	embStruct.print()

}

package main

import "fmt"

func main() {
	fmt.Println("Pointer")
	d := new(int)
	*d = 10
	fmt.Println(*d) //Output will be 10

	a := 10
	b := &a
	c := &b
	fmt.Println("Pointer", a)
	fmt.Println("Pointer", *b)
	fmt.Println("Pointer", *(*c))

	// method with pointer vs normal struct
	employePointer()
}

type employe struct {
	name string
	id   int
}

func employePointer() {
	emp := employe{name: "Nani", id: 1234}
	emp.details()

	emp1 := &employe{name: "Nani1", id: 12345}
	emp1.setName()
	fmt.Printf("Last Pointer Type %T and value name %v \n", emp1.name, emp1.name)
	fmt.Printf("Last Pointer Type %T and value id %v \n", emp1.id, emp1.id)
}

func (e employe) details() {
	fmt.Printf("Pointer Type %T and value name %v \n", e.name, e.name)
	fmt.Printf("Pointer Type %T and value id %v \n", e.id, e.id)
}

func (e *employe) setName() {
	e.name = "Jhon"
	fmt.Printf("Pointer Type %T and value name %v \n", e.name, e.name)
}

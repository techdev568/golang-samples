package main

import "fmt"

func main() {
	fmt.Print("Interface example")
	simpleInterface()
	interfaceExample()
}

type animal interface {
	breath()
	walk()
}

func simpleInterface() {
	var anim animal
	fmt.Print(anim)
}

//===================================
type lion struct {
	catogity string
	age      int
}

func interfaceExample() {
	//var anim animal

	anim := lion{catogity: "Lion", age: 100}
	anim.breath()
	anim.walk()
}

func (l lion) breath() {
	fmt.Print("Breath lion \n", l)
}

func (l lion) walk() {
	fmt.Print("walk lion \n", l)
}

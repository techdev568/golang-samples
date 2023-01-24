package main

import "fmt"

func main() {
	fmt.Println("Slice Declaration::")
	//  the size of a slice is flexible and can be changed.
	// slice with leanth 2 and capacity 2
	a := []int{3, 5, 7, 9, 10}
	fmt.Println("Slice: ", a)
	fmt.Println("length: ", len(a))
	fmt.Println("capacity: ", cap(a))

	//creating slice from another slice
	b := a[2:4]
	fmt.Println("Slice: ", b)
	// length of newly created slice = (end–start) :2
	fmt.Println("length: ", len(b))
	// capacity of newly created slice = (length_of_array–start) : 3
	fmt.Println("capacity: ", cap(b))
	// default value of second is length of array
	c := a[2:]
	// length of newly created slice = (end–start) :(default)5-2 = 3
	fmt.Println("length: ", len(c))
	// capacity of newly created slice = (length_of_array–start) : 5-2 = 3
	fmt.Println("capacity: ", cap(c))

	d := a[:3]
	// length of newly created slice = (end–start) :3-0(default) = 3
	fmt.Println("length: ", len(d))
	// capacity of newly created slice = (length_of_array–start) : 5-0(default) = 5
	fmt.Println("capacity: ", cap(d))
	fmt.Println("Slice using make:")
	sliceUsingMake()
	fmt.Println("Slice using New")
	sliceUsingNew()
	fmt.Println("Append Element")
	appendElementToSlice()
	fmt.Println("Append Slice to Slice")
	appendSliceToSlice()
	fmt.Println("Copy Slice")
	copySlice()
	fmt.Println("multiDimentionalSlice")
	multiDimentionalSlice()
}

func sliceUsingMake() {
	// length 3 and capacity 5
	a := make([]int, 3, 5)
	fmt.Println("length: ", len(a))
	fmt.Println("capacity: ", cap(a))
}

func sliceUsingNew() {
	c := make([]int, 3, 5)
	c = append(c, 4)
	a := new([]int)
	fmt.Println("length: ", len(*a))
	fmt.Println("capacity: ", cap(*a))
	b := new([]int)
	b = &c
	fmt.Println("length: ", b, len(*b))
	fmt.Println("capacity: ", cap(*b))
}

func appendElementToSlice() {
	numbers := make([]int, 3, 5)
	numbers[0] = 1
	numbers[1] = 2
	numbers = append(numbers, 4)
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}
func appendSliceToSlice() {
	numbers := make([]int, 3, 5)
	numbers1 := make([]int, 2, 5)
	numbers[0] = 1
	numbers[1] = 2
	numbers1[0] = 11
	numbers1[1] = 22
	// Notice '...' after the second slice. '...' is the operator
	// which means that the argument is a variadic parameter.
	// Meaning that during run time slice2 will be expanded to its individual elements which are passed as multiple arguments to the append function
	numbers = append(numbers, numbers1...)
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

func copySlice() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5)

	numbers := copy(dst, src)
	fmt.Printf("Number Of Elements Copied: %d\n", numbers)
	fmt.Printf("dst: %v\n", dst) //[1 2 3 4 5]
	fmt.Printf("src: %v\n", src) //[1 2 3 4 5]
}

func multiDimentionalSlice() {
	var twoDSlice = make([][]int, 2)
	twoDSlice[0] = []int{1, 2, 3}
	twoDSlice[1] = []int{4, 5, 6}
}

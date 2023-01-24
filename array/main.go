package main

import "fmt"

func main() {
	// An array is a contiguous collection of elements of the same type
	fmt.Println("Array Declaration::")
	arrayDeclaration()
	a := [5]string{"bb", "d", "f", "g", "kk"}
	fmt.Println("Array Iteration::")
	iterateArray(a)
	fmt.Println("Multi Dimentional Array")
	multiDimentionalArray()
}

func arrayDeclaration() {
	sample1 := [2]int{1, 2}
	fmt.Printf("Sample1: Len: %d, %v\n", len(sample1), sample1)

	//Only actual elements
	sample2 := [...]int{2, 3}
	fmt.Printf("Sample2: Len: %d, %v\n", len(sample2), sample2)

	//Only number of elements
	//'...' specifies the length of array will be equal to the number of actual elements
	sample3 := [2]int{}
	fmt.Printf("Sample3: Len: %d, %v\n", len(sample3), sample3)

	//Without both number of elements and actual elements
	//'...' specifies the length of array will be equal to the number of actual elements
	sample4 := [...]int{}
	fmt.Printf("Sample4: Len: %d, %v\n", len(sample4), sample4)
}

func iterateArray(letters [5]string) {
	for i, letter := range letters {
		fmt.Printf("%d %s\n", i, letter)
	}
}

func multiDimentionalArray() {
	sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i, row := range sample {
		fmt.Println("ROW :", i)
		for _, val := range row {
			fmt.Println(val)
		}
	}
}

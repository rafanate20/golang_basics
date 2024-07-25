package examples

import "fmt"

func UsingArrayAndSlices() {
	fmt.Printf("Chapter 006 - Using Arrays and Slices :\n")

	//vars
	var array [4]int
	fmt.Println("Arrays are mutable and golang initialize them:", array, len(array), cap(array))

	slice := []int{0, 1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("Slices are inmutable and you can add more values:", slice, len(slice), cap(slice))

	slice2 := []int{5, 6, 7}
	slice = append(slice, slice2...)
	fmt.Println("You can add another slice:", slice, len(slice), cap(slice))

	var slice3 []int = make([]int, 2, 5) //usign make to create a slice
	fmt.Printf("length:  %d - capacity %d - content:  %d \n", len(slice3), cap(slice3), slice3)
	slice3 = append(slice3, slice2...)
	fmt.Printf("length:  %d - capacity %d - content:  %d \n", len(slice3), cap(slice3), slice3)

	//slice de strings
	strSlice := []string{"Hello", "how", "are", "you"}

	for i, strVal := range strSlice {
		fmt.Print(i+1, "_", strVal, " ")
	}

	fmt.Println("\n----------------------------------------------------------------")
}

package main

import "fmt"

func arrays() {
	array := [3]int{1, 2, 3}
	fmt.Println(array)

	var names [4]string = [4]string{"Ram", "Shyam", "Mahesh", "Suresh"}
	fmt.Println(names)

	fractionals := [...]float64{2.5, 3.25, 8.9}
	fmt.Println(fractionals)

}

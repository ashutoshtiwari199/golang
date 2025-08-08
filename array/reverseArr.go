package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// temp := 0
	// for index := range a {
	// 	fmt.Println(index, len(a)-index)
	// 	fmt.Println(index, len(a)-index-1)
	// 	if index == len(a)-index {
	// 		fmt.Print("Broke")
	// 		break
	// 	} else {
	// 		temp = a[index]
	// 		a[index] = a[len(a)-1-index]
	// 		a[len(a)-index-1] = temp
	// 	}
	// }

	for i := 0; i <= len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}

	fmt.Println(a)
}

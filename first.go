package practice

import (
	"fmt"
)

func Sum(num int, size int) int {
	fmt.Println("enter number:")
	//var num int
	fmt.Scanln(&num)

	fmt.Println("enter size:")
	//var size int
	fmt.Scanln(&size)
	sum := 0 // another way for variable declaration ( declaration & intialization)

	for i := num; i <= size; i++ {
		sum = sum + i
	}
	fmt.Println("sum of ", size, "numbers:")
	return sum

}

func main() {
	var num, size int
	result := Sum(num, size)
	fmt.Println(result)

}

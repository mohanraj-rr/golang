package main

import "fmt"

func main() {

	var j int = 0
	var arr [5]string

	var i int = 1
	for i <= 5 {

		var choice int = 0
		fmt.Println("1-Enter list")
		fmt.Println("2-Display list")
		fmt.Println("3-Exit")
		fmt.Println("Enter your choice")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Enter task:")
			fmt.Scan(arr[j])
			j++
			continue
		}
		if choice == 2 {
			var k int = 0
			for k <= j {
				fmt.Println(arr[k])
				k++
			}
			continue
		}
		if choice == 3 {
			break
		}
		i++
	}

}

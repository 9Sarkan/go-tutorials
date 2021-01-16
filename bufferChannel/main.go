package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("there is no space to add!", i)
		}
	}

	for j := 0; j < counter+5; j++ {
		select {
		case number := <-numbers:
			fmt.Println(number)
		default:
			fmt.Println("No number!")
			break
		}
	}
}

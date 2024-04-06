package main

import "fmt"

func main() {
	fmt.Println("Welocome to simple calculator")
	input1 := ""
	for input1 != "exit" {
		fmt.Print(">")
		_, err := fmt.Scanf("%s", &input1)
		if err != nil {
			fmt.Println("sorry, invalid command")
		}
	}
}

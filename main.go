package main

import "fmt"

func main() {
	//if else
	//1. if condition {//code here}
	number := 10
	if number == 10 {
		fmt.Println("number = 10")
	}
	//2. if condition {code} else {code}
	if number < 10 {
		fmt.Println("number < 10")
	} else {
		fmt.Println("number = 10")
	}
	//3. if statement; condition { code }
	if a := 100; a > 100 {
		fmt.Println("a > 100")
	} else {
		fmt.Println("a = 100")
	}
}

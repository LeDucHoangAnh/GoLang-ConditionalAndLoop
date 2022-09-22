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
	//--switch - case: Không cần sử dụng break
	number1 := 5
	switch number1 {
	case 1:
		fmt.Println("number = 1")
	case 2:
		fmt.Println("number = 2")
	case 3:
		fmt.Println("number = 3")
	case 4:
		fmt.Println("number = 4")
	case 5, 6, 7, 8, 9:
		fmt.Println("number = 5,6,7,8,9")
	case 10:
		fmt.Println("number = 10")
	default:
		fmt.Println("Undefined")
	}
}

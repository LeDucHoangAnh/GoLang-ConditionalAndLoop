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
	//
	number2 := 10
	switch {
	case number2 > 10:
		fmt.Println("number > 10")
	case number2 == 10:
		fmt.Println("number = 10")
	}

	//FallThrough: Giúp chạy những case đằng sau khi bắt được case cần mà không dừng chương trình ngay
	number3 := 3
	switch number3 {
	case 1:
		fmt.Println("number = 1")
		fallthrough
	case 2:
		fmt.Println("number = 2")
		fallthrough

	case 3:
		fmt.Println("number = 3")
		fallthrough

	case 4:
		fmt.Println("number = 4")

	}

	//Break, goto
	number4 := 4
	switch number4 {
	case 1:
		fmt.Println("number = 1")
		fallthrough
	case 2:
		fmt.Println("number = 2")
		fallthrough

	case 3:
		if number4 == 10 {
			fmt.Println("Break")
			break
		}

	case 4:
		if number4 == 5 {
			goto handleNumberEqual10
		}
	handleNumberEqual10:
		fmt.Println("handle for case = 10")
	case 5:
		fmt.Println("number = 5")
	}
}

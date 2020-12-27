package main

import (
	"fmt"
	"strconv"
)

// jalankan  dengan mengganti variable x sesuka hati
// fungsi ini hanya menjalankan operator berurutan
// contoh jika x  =  10 -5 * 2 = 10
// yang akan di jalankan 10 - 2

// flow :
// dari fungsi 'main' akan memanggil fungsi  rekursif 'getResult'
// di fungsi 'getResult'    akan memanggil fungsi 'operation' untuk melakukan operasi matematika antara bilangan 1 dan bilangan 2
// di fungsi 'getResult'    akan memanggil fungsi 'getSlice' untuk melakukan slice array yang berisi interger / int
// di fungsi 'getResult'    akan memanggil fungsi 'getSliceOperator' untuk melakukan slice array yang berisi string

func main() {
	x := "10 -5 *2 "

	getOperation(x)
}

func getOperation(input string) {
	operator := []string{
		"+",
		"%",
		"/",
		"*",
		"-",
	}

	var currentOperator []string
	number := []int{}
	for _, item := range input {
		char := fmt.Sprintf("%c", item)
		for _, op := range operator {

			if char == op {
				if len(currentOperator) == 0 {
					currentOperator = append(currentOperator, "+")

				}
				currentOperator = append(currentOperator, op)
			}
		}

		temp := 0
		if char == "0" {
			number = append(number, 0)
		} else {
			temp, _ = strconv.Atoi(char)
			if temp != 0 {

				number = append(number, temp)
			}
		}

	}
	// fmt.Println(currentOperator)
	result := getResult(number, currentOperator, 0)
	fmt.Println(result)

}

func getResult(number []int, currentOperator []string, result int) int {

	indexOperator := (len(currentOperator) - len(number))

	indexOperator -= 1
	if indexOperator < 0 {
		indexOperator = 0
	}

	if len(number) == 1 {

		return operation(result, number[0], currentOperator[indexOperator])
	} else {

		return getResult(
			getSlice(number),
			getSliceOperator(currentOperator),
			operation(result,
				number[0], currentOperator[indexOperator]))

	}

}

func operation(number1 int, number2 int, currentOperator string) int {
	var result int
	if currentOperator == "+" {
		result = number1 + number2
	} else if currentOperator == "%" {
		result = number1 % number2
	} else if currentOperator == "/" {
		result = number1 / number2
	} else if currentOperator == "*" {
		result = number1 * number2

	} else if currentOperator == "-" {

		result = number1 - number2
	}
	// fmt.Println(number1, number2, "===", result)
	return result
}

func getSlice(number []int) []int {

	if len(number) == 1 {
		return number
	} else {

		var result []int
		for i := range number {

			if i > 0 {

				result = append(result, number[i])
			}
		}
		return result
	}

}
func getSliceOperator(operator []string) []string {

	if len(operator) == 1 {
		return operator
	} else {

		var result []string
		for i := range operator {

			if i > 0 {

				result = append(result, operator[i])
			}
		}
		return result
	}

}

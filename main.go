package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getRomanVal(roman string) int {

	var romanMap = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	return romanMap[roman]
}

func convertToRoman(x int) string {
	romanNSlice := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	romanDSlice := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}

	if x <= 0 {
		panic("Error: Unable to execute")
	}

	if x >= 10 {
		var result string
		var decPart = romanDSlice[x/10-1]
		if x%10-1 >= 0 {
			var modPart = romanNSlice[x%10-1]
			result = string(decPart) + string(modPart)
		} else {
			result = string(decPart)
		}

		return result
	} else {
		return romanNSlice[x-1]
	}
}

func sum(x, y int) (z int) {
	z = x + y
	return
}

func subtraction(x, y int) (z int) {
	z = x - y
	return
}

func multiplications(x, y int) (z int) {
	z = x * y
	return
}

func division(x, y int) (z int) {
	z = x / y
	return
}

func getResult(operation string) func(int, int) int {

	if operation == "*" {
		return multiplications
	} else if operation == "/" {
		return division
	} else if operation == "-" {
		return subtraction
	} else {
		return sum
	}
}

func validData(expression string) bool {
	var validIDRoman = regexp.MustCompile(`^(X|IX|IV|V?I{0,3})([-+*\/]{1})(X|IX|IV|V?I{0,3})$`)
	var validID = regexp.MustCompile(`^([1-9]|10)([-+*\/]{1})([1-9]|10)$`)

	return validIDRoman.MatchString(expression) || validID.MatchString(expression)
}

func main() {

	var inputData string
	fmt.Scan(&inputData)
	inputData = strings.ReplaceAll(inputData, " ", "")

	if validData(inputData) {
		var operationIndex int = strings.IndexAny(inputData, "-+*/")
		var operationSymb string = inputData[operationIndex : operationIndex+1]

		firstNum, fErr := strconv.Atoi(inputData[0:operationIndex])
		secondNum, sErr := strconv.Atoi(inputData[operationIndex+1:])

		result := getResult(operationSymb)

		if fErr != nil && sErr != nil {

			firstNum = getRomanVal(inputData[0:operationIndex])
			secondNum = getRomanVal(inputData[operationIndex+1:])

			fmt.Println(convertToRoman(result(firstNum, secondNum)))

		} else {
			fmt.Println(result(firstNum, secondNum))
		}

	} else {
		panic("Error: Incorrect data")
	}
}

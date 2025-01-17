package string_sum

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if (input == "") {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	input = strings.ReplaceAll(input, " ", "")
	isPlus := false
	stringNumbers := [2]string{}
	splittedString := strings.Split(input, "+")
	if (len(splittedString) != 2) {
		splittedString := strings.Split(input, "-")
		if (len(splittedString) == 2) {
			stringNumbers[0] = string(splittedString[0])
			stringNumbers[1] = string(splittedString[1])
		} else if (len(splittedString) != 2 && splittedString[0] == "") {
			stringNumbers[0] = "-" + string(splittedString[1])
			stringNumbers[1] = string(splittedString[2])
		} else {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
	} else {
		isPlus = true
		stringNumbers[0] = string(splittedString[0])
		stringNumbers[1] = string(splittedString[1])
	}
	var parsedFirstNumber, parsedFirstNumberErr = strconv.Atoi(stringNumbers[0])
	if (parsedFirstNumberErr != nil) {
		return "", fmt.Errorf("%w", parsedFirstNumberErr)
	}
	var parsedSecondNumber, parsedSecondNumberErr = strconv.Atoi(stringNumbers[1])
	if (parsedSecondNumberErr != nil) {
		return "", fmt.Errorf("%w", parsedSecondNumberErr)
	}
	if (isPlus) {
		return fmt.Sprint(parsedFirstNumber + parsedSecondNumber), nil
	} else {
		return fmt.Sprint(parsedFirstNumber - parsedSecondNumber), nil
	}
}

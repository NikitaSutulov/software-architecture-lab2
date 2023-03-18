package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// remove element with index
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// function to detect if the string is actually a number in a string form
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// function to check if there is a string in an array
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// PrefixToInfix converts a prefix expression to an infix expression
// Input: a string representing a prefix expression
// Output: a string representing the equivalent infix expression
// Error: returns an error if unable to convert the input expression
func PrefixToInfix(input string) (string, error) {
	inputSplit := strings.Split(input, " ")
	var operators = []string{"+", "-", "*", "/", "^"}
	var s = []string{}

	if len(inputSplit) < 3 {
		return "", fmt.Errorf("Unable to convert")
	} else if len(inputSplit) == 3 {
		var operator = inputSplit[0]
		var left = inputSplit[1]
		var right = inputSplit[2]
		if contains(operators, operator) && isNumeric(left) && isNumeric(right) {
			return fmt.Sprintf("%s %s %s", left, operator, right), nil
		} else {
			return "", fmt.Errorf("Unable to convert")
		}
	}

	for i := len(inputSplit) - 1; i >= 0; i-- {
		if isNumeric(inputSplit[i]) {
			s = append(s, inputSplit[i])
		} else {
			var op1 = s[len(s)-2]
			var op2 = s[len(s)-1]
			s = remove(s, len(s)-1)
			s = remove(s, len(s)-1)
			if len(s) > 0 {
				s = remove(s, len(s)-1)
			}
			if contains(operators, inputSplit[i]) {
				s = append(s, fmt.Sprintf("(%s %s %s)", op2, inputSplit[i], op1))
			} else {
				return "", fmt.Errorf("Unable to convert")
			}
		}
	}

	return s[len(s)-1], nil

}

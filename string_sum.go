package string_sum

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	input = strings.TrimSpace(input)
	if input == "" {
		log.Fatal(errorEmptyInput)
	}

	el := strings.Split(input, "")
	var el2 []string
	for i := 0; i < len(el); i++ {
		if el[i] == " " {
			continue
		}
		if el[i] == "+" {
			el[i] = ","
			el2 = append(el2, el[i])
			continue
		}
		/* 		if el[i] == "-" {
			el[i] = ",-"
		} */
		el2 = append(el2, el[i])
	}
	input = strings.Join(el2, "")
	el = strings.Split(strings.TrimSpace(input), ",")
	if len(el) > 2 {
		log.Fatal(errorNotTwoOperands)
	}
	fmt.Println("el=", el, len(el))
	if len(el) < 2 {
		log.Fatal(errorNotTwoOperands)
	}

	a, err := strconv.ParseInt(el[0], 10, 64)
	if err != nil {
		log.Fatal("Error summand number one is not correct", err)
	}

	b, err := strconv.ParseInt(el[1], 10, 64)
	if err != nil {
		log.Fatal("Error summand number two is not correct", err)
	}
	output = fmt.Sprint(strconv.FormatInt(a+b, 10))
	return output, nil
}

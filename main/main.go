package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func main() {
	z := "              \n\t55 +                -3"
	fmt.Println(StringSum(z))
}

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

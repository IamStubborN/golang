package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

var cardNumber = "5168742718319332"

func main() {
	result, err := Validate(cardNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

// Validate validate string by Luhn algoritm
func Validate(str string) (string, error) {
	err := runSteps(&str, trimSpaces, checkLengh, calculateResults)
	if err != nil {
		return "", err
	}
	return "You have valid card number", err
}

func runSteps(str *string, funcs ...func(*string) error) error {
	var err error
	for _, fn := range funcs {
		err = fn(str)
		if err != nil {
			break
		}
	}
	return err
}

func checkLengh(str *string) error {
	if len(*str) <= 1 {
		return fmt.Errorf("You have bad lengh %v", len(*str))
	}
	return nil
}

func trimSpaces(str *string) error {
	tempString := ""
	for _, r := range *str {
		if unicode.IsSpace(r) {
			continue
		} else if unicode.IsDigit(r) {
			tempString += string(r)
		} else {
			return fmt.Errorf("You have bad symbol %q", string(r))
		}
	}
	*str = tempString
	return nil
}

func calculateResults(str *string) error {
	var sum = 0
	for i, r := range *str {
		if i%2 == 0 {
			number, err := strconv.Atoi(string(r))
			if err != nil {
				return err
			}
			calcNumber := number * 2
			if calcNumber > 9 {
				calcNumber -= 9
			}
			sum += calcNumber
		} else {
			number, err := strconv.Atoi(string(r))
			if err != nil {
				return err
			}
			sum += number
		}
	}
	if sum%10 == 0 {
		return nil
	}
	return fmt.Errorf("You have not valid card with %v sum", sum)
}

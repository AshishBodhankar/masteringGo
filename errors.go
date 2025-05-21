package main

import (
	"errors"
	"fmt"
)

//	type divideError struct {
//		dividend float64
//	}
//
//	func (d divideError) Error() string {
//		return fmt.Sprintf("cannot divide %v by zero", d.dividend)
//	}
func validateStatus(status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	} else if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	} else {
		return nil
	}
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		//return 0, divideError{dividend: dividend}
		return 0, errors.New("no dividing by 0")
	}
	return dividend / divisor, nil
}

/* ----------------------------------------------------------------------------------------------  */
func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	customerCost, customerError := sendSMS(msgToCustomer)
	spouseCost, spouseError := sendSMS(msgToSpouse)

	if customerError != nil {
		return 0, customerError
	} else if spouseError != nil {
		return 0, spouseError
	} else {
		return customerCost + spouseCost, nil
	}
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("Can't send texts over %d characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func main() {
	msgToCustomer := "Wassup bro"
	msgToSpouse := "wassup sister"
	fmt.Println(sendSMSToCouple(msgToCustomer, msgToSpouse))
	fmt.Println(divide(20, 0))
	fmt.Println(validateStatus("Hi bro wassup!"))
}

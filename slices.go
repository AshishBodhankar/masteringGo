package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

func getMessageCosts(messages []string) []float64 {

	// func make([]T, len, cap) []T
	costs := make([]float64, len(messages))
	for i := 0; i < len(costs); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planFree:
		return messages[:], nil
	case planPro:
		return messages[:2], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	var costs [3]int
	cost := 0
	for i := 0; i < len(messages); i++ {
		cost += len(messages[i])
		costs[i] = cost
	}
	return messages, costs
}

func main() {

	fmt.Println(getMessageWithRetries("Ashish", "Vaishnavi", "Vaishish"))
}

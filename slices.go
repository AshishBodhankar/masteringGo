package main

import (
	"errors"
	"fmt"
)

/* SYNTAX for RANGE in Golang

for INDEX, ELEMENT := range SLICE {
}

*/

func indexOfFirstBadWord(msg []string, badWords []string) (firstIndex int) {
	firstIndex = -1

	for index, message := range msg {
		for _, badWord := range badWords {
			if message == badWord {
				return index
			}
		}
	}
	return
}

/*--------------------------------------------------------------------------------------------------------------------*/

// Custom defined Append() function

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

/*

BUILT-IN append function

func append(slice []Type, elems ...Type) []Type

BELOW ARE ALL VALID because append() is a variadic function:
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
*/

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	costValues := make([]float64, 0)
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			costValues = append(costValues, costs[i].value)
		}
	}
	return costValues
}

/*
variadic function receives variadic arguments as a slice.
Therefore, use for loop to read/select arguments iteratively from the slice.

The spread operator allows us to pass a slice into a variadic function.
The spread operator consists of three dots following the slice in the function call.
*/

func sum(nums ...int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

// Slices in Go are of variable length. Slices reference a fixed sized array in the background
func getMessageCosts(messages []string) []float64 {

	// func make([]T, len, cap) []T
	costs := make([]float64, len(messages))
	for i := 0; i < len(costs); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planPro:
		return messages[:], nil
	case planFree:
		return messages[:2], nil
	default:
		return nil, errors.New("unsupported plan")
	}
}

// ARRAYS in Go are fixed length group of elements of the same type
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

	names := []string{"Ashish", "Vaishu", "Aie"}
	printStrings(names...)
	printStrings("Ashish", "Vaishu", "Aie")

	nums := []int{1, 24, 6}
	fmt.Println(sum(nums...))
	fmt.Println(sum(1, 24, 6))

	cost1 := cost{day: 1, value: 2}
	cost2 := cost{day: 2, value: 223}
	cost3 := cost{day: 2, value: 12}
	cost4 := cost{day: 5, value: 13}
	costs := []cost{cost1, cost2, cost3, cost4}
	fmt.Println(getDayCosts(costs, 2))
}

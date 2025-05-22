package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {

	for index, message := range messages {
		messages[index].tags = tagger(message)
	}
	return messages
}

func tagger(msg sms) []string {

	tags := []string{}

	content := strings.ToLower(msg.content)

	if strings.Contains(content, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(content, "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}

/*--------------------------------------------------------------------------------------------------------------------*/
func isValidPassword(password string) (bool, error) {
	if len(password) < 5 || len(password) >= 12 {
		return false, errors.New("password length should be at least 5 characters long but no more " +
			"than 12 characters")
	}
	oneUpper := false
	isDigit := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			oneUpper = true
		}
		if unicode.IsDigit(char) {
			isDigit = true
		}
	}
	if !oneUpper {
		return false, errors.New("password should contain at least one uppercase letter")
	}
	if !isDigit {
		return false, errors.New("password should contain at least one digit")
	}
	return true, nil
}

/*--------------------------------------------------------------------------------------------------------------------*/

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	result := []Message{}
	for _, message := range messages {
		if message.Type() == filterType {
			result = append(result, message)
		}
	}
	return result
}

/*--------------------------------------------------------------------------------------------------------------------*/

// Slice of Slices

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for r := 0; r < rows; r++ {
		row := []int{}
		for c := 0; c < cols; c++ {
			row = append(row, r*c)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

/*--------------------------------------------------------------------------------------------------------------------*/

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

	//costValues := make([]float64, 0)
	costValues := []float64{}

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

	matrix := createMatrix(5, 10)
	fmt.Println(matrix)

	tm := TextMessage{Sender: "Ashish", Content: "wassup my nigga!"}
	mm := MediaMessage{Sender: "Vaishu", MediaType: ".mp4", Content: "sare Jahan se Acha! Hindustan hamara"}
	lm := LinkMessage{Sender: "Aie", URL: "http://www.koogle.com", Content: "Nothing important"}
	messages := []Message{tm, mm, lm}
	fmt.Println(filterMessages(messages, "media"))

	fmt.Println(isValidPassword("Escanor#123"))

	messages2 := []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
		// Additional messages...
	}
	taggedMessages := tagMessages(messages2, tagger)
	fmt.Println(taggedMessages)
}

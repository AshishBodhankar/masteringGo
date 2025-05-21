package main

import (
	"fmt"
	//"math"
)

/* ----------------------------------------------------------------------------------------------  */
//func getExpenseReport(e expense) (string, float64) {
//	email_, isEmail := e.(email)
//	if isEmail {
//		return email_.format(), email_.cost()
//	} else if sms_, isSMS := e.(sms); isSMS {
//		return sms_.toPhoneNumber, sms_.cost()
//	} else {
//		return "", e.cost()
//	}
//}
func getExpenseReport(e Expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.format(), v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", v.cost()
	}
}

type Expense interface {
	cost() float64
}

type invalid struct{}

func (i invalid) cost() float64 {
	return 0.0
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	} else {
		return 0.01 * float64(len(e.body))
	}
}
func (e email) format() string {
	if e.isSubscribed == false {
		return fmt.Sprintf("'%s' | Not Subscribed | %s", e.body, e.toAddress)
	} else {
		return fmt.Sprintf("'%s' | Subscribed | %s", e.body, e.toAddress)
	}
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	} else {
		return float64(len(s.body)) * 0.03
	}
}

func SendMessage(f Formatter) string {
	return f.format()
}

type Formatter interface {
	format() string
}

type PlainText struct {
	message string
}

func (p PlainText) format() string {
	return fmt.Sprintf(p.message)
}

type Bold struct {
	message string
}

func (b Bold) format() string {
	return fmt.Sprintf("**%s**", b.message)
}

type Code struct {
	message string
}

func (c Code) format() string {
	return fmt.Sprintf(`%s`, c.message)
}

/* ----------------------------------------------------------------------------------------------  */
type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	} else {
		return d.priorityLevel
	}
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (s systemAlert) importance() int {
	return 100
}

// ?

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, v.importance()
	case groupMessage:
		return v.groupName, v.importance()
	case systemAlert:
		return v.alertCode, v.importance()
	default:
		return "", 0
	}
}

/* ----------------------------------------------------------------------------------------------  */

func test(e employee) string {
	return fmt.Sprintf("%s has a salary equal to: %d", e.getName(), e.getSalary())
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (f fullTime) getName() string {
	return f.name
}
func (f fullTime) getSalary() int {
	return f.salary
}

/* ----------------------------------------------------------------------------------------------  */

func main() {
	fmt.Println(test(contractor{name: "Ashish", hourlyPay: 60, hoursPerYear: 1920}))
	fmt.Println(test(fullTime{name: "Vaishnavi", salary: 150000}))

	_sms := sms{isSubscribed: true, body: "Hi Vaishu em chestunnavu", toPhoneNumber: "5716397308"}
	//_invalid := invalid{}
	_email := email{isSubscribed: false, body: "Hi Vaishu, This is Ashish. What are you upto?", toAddress: "Jersey City, NJ"}
	fmt.Println(getExpenseReport(_sms))
	fmt.Println(getExpenseReport(_email))
	fmt.Println(SendMessage(Code{message: "Ashish"}))
}

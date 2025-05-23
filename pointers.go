package main

import (
	"errors"
	"fmt"
	"strings"
	//"unicode"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(cstPtr *customer, trn transaction) error {
	if trn.customerID != (*cstPtr).id {
		return errors.New("customer ID is not matching. Can't process updation!")
	} else if (trn.transactionType != transactionDeposit) && (trn.transactionType != transactionWithdrawal) {
		return errors.New("unknown transaction type")
	} else if (trn.transactionType == transactionWithdrawal) && (trn.amount > (*cstPtr).balance) {
		return errors.New("insufficient funds")
	} else {
		if trn.transactionType == transactionDeposit {
			(*cstPtr).balance += trn.amount
		} else if trn.transactionType == transactionWithdrawal {
			(*cstPtr).balance -= trn.amount
		}
		return nil
	}
}

/*--------------------------------------- Pointer Receivers ---------------------------------------------------*/
type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type car struct {
	color string
}

func (c *car) setColor(color string) {
	(*c).color = color
}

//func (c car) setColor(color string) {
//	c.color = color
//}

/*-------------------------------------------------------------------------------------------------------------*/

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(aPtr *Analytics, mPtr *Message) {
	if (*mPtr).Success {
		(*aPtr).MessagesTotal++
		(*aPtr).MessagesSucceeded++
	} else {
		(*aPtr).MessagesTotal++
		(*aPtr).MessagesFailed++
	}
}

/*-------------------------------------------------------------------------------------------------------------*/

func removeProfanity(message *string) {

	if message == nil {
		return // Avoid panic by returning early
	}
	words := map[string]string{"fubb": "****", "shiz": "****", "witch": "*****"}
	for word, replacement := range words {
		*message = strings.ReplaceAll(*message, word, replacement)
	}

}

func main() {
	test1 := "English, motherfubber, do you speak it?"
	removeProfanity(&test1)

	words := map[string]string{"fubb": "****", "shiz": "****", "witch": "*****"}
	fmt.Println(test1)
	fmt.Println(words)

	newMessage := Message{Recipient: "mickey", Success: true}
	initialAnalytics := Analytics{MessagesTotal: 0, MessagesFailed: 0, MessagesSucceeded: 0}
	getMessageText(&initialAnalytics, &newMessage)
	fmt.Println(initialAnalytics)

	newMessage2 := Message{Recipient: "minnie", Success: false}
	initialAnalytics2 := Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1}
	getMessageText(&initialAnalytics2, &newMessage2)
	fmt.Println(initialAnalytics2)

	c := car{color: "white"}
	c.setColor("blue")
	fmt.Println(c.color)

	cust1 := customer{id: 1, balance: 100.0}
	cust2 := customer{id: 2, balance: 200.0}
	trn1 := transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit}
	trn2 := transaction{customerID: 2, amount: 100.0, transactionType: transactionWithdrawal}
	trn3 := transaction{customerID: 2, amount: 150.0, transactionType: transactionDeposit}
	fmt.Println(updateBalance(&cust1, trn1))
	fmt.Println(updateBalance(&cust2, trn2))
	fmt.Println(updateBalance(&cust2, trn3))
	fmt.Println(cust1)
	fmt.Println(cust2)
}

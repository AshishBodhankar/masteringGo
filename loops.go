package main

import (
	"fmt"
)

/*
for INITIAL; CONDITION; AFTER{
  // do something
}
*/

func countConnections(groupSize int) int {
	connections := 0
	for i := 0; i < groupSize; i++ {
		for j := i + 1; j < groupSize; j++ {
			connections += 1
		}
	}
	return connections
}

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func fizzbuzz() {
	i := 1
	for i <= 100 {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance > 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func bulkSend(numMessages int) float64 {
	cost := 0.0
	for i := 0; i < numMessages; i++ {
		cost += 1.0 + (0.01 * float64(i))
	}
	return cost
}

func maxMessages(thresh int) int {
	cost := 0
	for i := 0; ; i++ {
		cost += 100 + i
		if cost > thresh {
			return i
		}
	}
}

func main() {
	fmt.Println(countConnections(5))
	//printPrimes(100)
	//fizzbuzz()
	//fmt.Println(bulkSend(10))
	//fmt.Println(maxMessages(1500))
}

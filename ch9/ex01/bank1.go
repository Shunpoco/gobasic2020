package main

import "fmt"

var deposits = make(chan int) // 入金額の送信
var balances = make(chan int) // 残高の受信
var drawBool = make(chan bool)
var drawAmount = make(chan int)

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func Withdraw(amount int) bool {
	drawAmount <- amount
	return <-drawBool
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case draw := <-drawAmount:
			if balance < draw {
				drawBool <- false
			} else {
				balance -= draw
				drawBool <- true
			}

		}
	}
}

func init() {
	go teller()
}

func main() {
	Deposit(1000)
	fmt.Println(Balance())
	fmt.Println(Withdraw(1000))
	fmt.Println(Balance())
	fmt.Println(Withdraw(1000))
}

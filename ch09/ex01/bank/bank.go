package bank

var deposits = make(chan int)
var balances = make(chan int)

var withdraws = make(chan int)
var ok = make(chan bool)

func Deposit(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	withdraws <- amount
	return <-ok
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraws:
			ok <- balance >= amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

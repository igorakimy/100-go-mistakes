package main

type Customer struct {
	balance int
}

func (c *Customer) Balance() int {
	return c.balance
}

func (c *Customer) SetBalance(amount int) {
	c.balance = amount
}

func main() {
	customer := Customer{
		balance: 100,
	}

	// Полезный геттер
	currentBalance := customer.Balance()
	if currentBalance < 0 {
		// Полезный сеттер
		customer.SetBalance(0)
	}
}

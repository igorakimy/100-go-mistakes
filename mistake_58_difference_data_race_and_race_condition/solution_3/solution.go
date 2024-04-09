package main

func main() {
	i := 0
	ch := make(chan int)

	go func() {
		// Уведомление горутины об увеличении на 1.
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	// Увеличение i от того её значения, которое было получено из канала.
	i += <-ch
	i += <-ch
}

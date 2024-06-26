package main

// гонка данных (data race) - происходит, когда две или более горутины
// одновременно обращаются к одной и той же ячейке памяти и по крайней
// мере одна из них выполняет запись в эту ячейку.

func main() {
	i := 0

	go func() {
		// увеличить i на 1.
		i++
	}()

	go func() {
		i++
	}()
}

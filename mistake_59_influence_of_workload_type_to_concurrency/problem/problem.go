package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func read(r io.Reader) (int, error) {
	count := 0
	for {
		// Чтение 1024 байт.
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			// Остановка цикла, когда достигнут его конец.
			if err == io.EOF {
				break
			}
			return 0, err
		}
		// Увеличение count на величину, получающуюся на основе
		// результата функции task. task выполняется последовательно,
		// но есть возможность сделать её выполнение параллельным.
		count += task(b)
	}
	return count, nil
}

func task(b []byte) int {
	// Выполнение какой-то задачи...
	return 1
}

func main() {
	file := "mistake_59_influence_of_workload_type_to_concurrency/content.txt"

	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(b)

	count, err := read(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

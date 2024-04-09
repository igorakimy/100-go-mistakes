package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

func read(r io.Reader) (int, error) {
	var count int64
	wg := sync.WaitGroup{}
	// Определение размера пула горутин.
	var n = 10

	// Создание канала с емкостью, равную пулу.
	ch := make(chan []byte, n)

	// Добавление n к WaitGroup.
	wg.Add(n)

	// Создание пула из n горутин.
	for i := 0; i < n; i++ {
		go func() {
			// Вызов метода Done, когда горутина получила данные из канала.
			defer wg.Done()
			// Каждая горутина получает данные из общего канала.
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	for {
		b := make([]byte, 1024)
		// Чтение из читателя r в слайс байт b.
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		// Публикация новой задачи в канале после каждого чтения.
		ch <- b
	}

	close(ch)
	// Ожидание завершения группы Wait перед возвратом.
	wg.Wait()

	return int(count), nil
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

package main

import (
	"fmt"
	"sync"
)

// Состояние гонки (race condition) - возникает, когда поведение зависит от
// последовательности или времени выполнения событий, которые невозможно
// контролировать.

func main() {
	i := 0
	mutex := sync.Mutex{}

	// В данном примере произойдет состояние гонки, т.к. порядок выполнения
	// горутин является случайным, то в i будет записана либо 1, либо 2.

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		// Первая горутина присваивает переменной i значение, равное 1.
		i = 1
	}()

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		// Вторая горутина присваивает переменной i значение, равное 2.
		i = 2
	}()

	fmt.Println(i)
}
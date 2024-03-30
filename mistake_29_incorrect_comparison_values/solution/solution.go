package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	id         string
	operations []float64
}

func (a customer) equal(b customer) bool {
	// Проводится сравнение полей id.
	if a.id != b.id {
		return false
	}
	// Проверяется длина обоих срезов.
	if len(a.operations) != len(b.operations) {
		return false
	}
	// Проводится сравнение каждого элемента обоих срезов.
	for i := 0; i < len(a.operations); i++ {
		if a.operations[i] != b.operations[i] {
			return false
		}
	}
	return true
}

func main() {
	cust1 := customer{id: "x", operations: []float64{1.}}
	cust2 := customer{id: "x", operations: []float64{1.}}

	// Первый способ сравнения (медленный)
	fmt.Println(reflect.DeepEqual(cust1, cust2)) // true

	// Второй способ сравнения (более быстрый)
	fmt.Println(cust1.equal(cust2))
}

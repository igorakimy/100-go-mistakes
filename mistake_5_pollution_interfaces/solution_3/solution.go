package main

import "fmt"

type intConfigGetter interface {
	Get() int
}

type IntConfig struct {
	// ...
}

func (c *IntConfig) Get() int {
	// Получить конфигурацию
	return 0
}
func (c *IntConfig) Set(value int) {
	// Обновить конфигурацию
}

type Foo struct {
	threshold intConfigGetter
}

func (f Foo) Bar() {
	// Чтение конфигурации
	threshold := f.threshold.Get()
	fmt.Println(threshold)
	// ...
}

// Вводится геттер конфигурации

func NewFoo(threshold intConfigGetter) Foo {
	return Foo{threshold: threshold}
}

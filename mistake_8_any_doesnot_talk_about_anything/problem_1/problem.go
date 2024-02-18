package main

func main() {
	var i any

	// Тип int
	i = 42
	// Тип string
	i = "foo"
	// Структура
	i = struct {
		s string
	}{
		s: "bar",
	}
	// Функция
	i = f
	// Присвоение значения пустому идентификатору
	_ = i
}

func f() {}

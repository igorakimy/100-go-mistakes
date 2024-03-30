package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 1_000_000
	// Создаем пустую карту с указателями,
	// что значительно сэкономит потребление памяти.
	m := make(map[int]*[128]byte)
	printAlloc()

	// Добавление 1 миллиона элементов в карту.
	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	printAlloc()

	// Удаление 1 миллиона элементов из карты.
	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// Ручной запуск сборщика мусора.
	runtime.GC()
	printAlloc()
	// Сохранение ссылки на m, чтобы сборщик мусора не "чистил" карту.
	runtime.KeepAlive(m)
}

func randBytes() *[128]byte {
	return &[128]byte{42, 180, 41, 233, 106, 182, 230, 253, 108, 222, 77, 27, 106, 249, 138, 124, 229, 108, 226, 113, 141, 227, 83, 173, 152, 182, 244, 188, 129, 221, 178, 162, 218, 153, 244, 237, 57, 239, 141, 41, 176, 39, 214, 181, 240, 34, 124, 11, 27, 16, 41, 101, 216, 210, 27, 175, 99, 114, 130, 169, 208, 195, 84, 46, 142, 57, 230, 242, 103, 18, 118, 157, 88, 33, 103, 40, 221, 64, 46, 3, 47, 213, 46, 61, 12, 166, 113, 93, 37, 110, 145, 222, 255, 29, 54, 188, 40, 43, 55, 175, 206, 16, 1, 255, 103, 158, 34, 75, 141, 58, 27, 173, 245, 223, 10, 69, 5, 196, 45, 183, 105, 31, 98, 146, 100, 87, 28, 21}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

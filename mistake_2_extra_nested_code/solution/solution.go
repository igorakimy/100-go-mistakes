package main

import (
	"errors"
	"fmt"
)

func join(s1, s2 string, max int) (string, error) {
	// Выравняйте (happy path) по левому краю — так вы сможете
	// быстро просмотреть, что происходит ниже на каком-то одном
	// уровне и увидеть, что на нем ожидаемо выполняется
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return s1 + s2, nil
}

func main() {
	s, err := join("test1", "test2", 12)
	if err != nil {
		return
	}
	fmt.Println(s)
}

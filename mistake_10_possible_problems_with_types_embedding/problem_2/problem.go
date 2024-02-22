package main

import (
	"io"
	"os"
)

// Избыточное определение методов. Вместо этого стоит
// воспользоваться встраиванием

type Logger struct {
	writeCloser io.WriteCloser
}

func (l Logger) Write(p []byte) (int, error) {
	return l.writeCloser.Write(p)
}

func (l Logger) Close() error {
	return l.writeCloser.Close()
}

func main() {
	l := Logger{writeCloser: os.Stdout}
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}

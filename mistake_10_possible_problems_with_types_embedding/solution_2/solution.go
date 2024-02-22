package main

import (
	"io"
	"os"
)

type Logger struct {
	io.WriteCloser
}

func main() {
	l := Logger{WriteCloser: os.Stderr}
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}

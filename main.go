package main

import (
	"os"

	"github.com/Krishanu230/Flipbook-Language/repl"
)

func main() {
	fname := string(os.Args[1])
	fileHandle, _ := os.Open(fname)
	defer fileHandle.Close()
	repl.Start(fileHandle, os.Stdout)
}

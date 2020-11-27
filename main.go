package main

import (
	"Flipbook/repl"
	"os"
)

func main() {
	fname := string(os.Args[1])
	fileHandle, _ := os.Open(fname)
	defer fileHandle.Close()
	repl.Start(fileHandle, os.Stdout)
}

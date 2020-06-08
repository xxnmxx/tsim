package main

import (
	"os"

	"github.com/xxnmxx/tsim"
)

func main() {
	tsim.Start(os.Stdin, os.Stdout)
}

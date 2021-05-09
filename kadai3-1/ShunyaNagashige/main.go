package main

import (
	"os"

	"github.com/ShunyaNagashige/golang-de/typing"
)

func main() {
	ch := typing.Input(os.Stdin)

	typing.Start(ch)
}

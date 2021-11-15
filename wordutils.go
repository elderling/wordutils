package main

import (
	"fmt"

	"github.com/elderling/findwords"
)

func main() {
	words := findwords.NewWordList()

	fmt.Println(words)
}

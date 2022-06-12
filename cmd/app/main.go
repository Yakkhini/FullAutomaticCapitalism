package main

import (
	"fmt"

	"fultocapital/internal/pull"
)

func main() {
	fmt.Println("Hello, World!")

	// print market response
	data := pull.Pull()
    fmt.Println(data)
}

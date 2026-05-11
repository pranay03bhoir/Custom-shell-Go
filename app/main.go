package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	scanner := bufio.NewScanner(os.Stdin)
	var userInput any
	userInput = scanner.Scan()
	fmt.Print("$ ")
	if userInput != nil {
		fmt.Println(userInput, ": command not found")
	}
}

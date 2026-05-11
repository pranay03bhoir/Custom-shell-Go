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
	fmt.Print("$ ")
	userInput = scanner.Scan()
	if userInput != nil {
		fmt.Printf("%s: command not found\n", scanner.Text())
	}
}

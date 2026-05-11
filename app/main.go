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
	for {
		fmt.Print("$ ")
		exit := scanner.Text()
		if exit == "exit" {
			return
		}
		if scanner.Scan() {
			if scanner.Text() == "exit" {
				return
			}
			fmt.Printf("%s: command not found\n", scanner.Text())
		}
	}
}

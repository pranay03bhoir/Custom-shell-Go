package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
			command := scanner.Text()
			command = strings.TrimSpace(command)
			if command == "exit" {
				return
			} else if strings.HasPrefix(command, "echo ") {
				fmt.Println(command[5:])
			} else if strings.HasPrefix(command, "type ") {
				extractBuiltInCommand := command[5:]
				extractBuiltInCommand = strings.TrimSpace(extractBuiltInCommand)
				switch extractBuiltInCommand {
				case "echo":
					fmt.Println(extractBuiltInCommand, "is a shell builtin")
				case "exit":
					fmt.Println(extractBuiltInCommand, "is a shell builtin")
				case "type":
					fmt.Println(extractBuiltInCommand, "is a shell builtin")
				default:
					fmt.Printf("%s: not found\n", extractBuiltInCommand)
				}
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}
}

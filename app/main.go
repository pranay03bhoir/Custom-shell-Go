package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
					foundPath := findFilePath(extractBuiltInCommand)
					if foundPath != "" {
						fmt.Println(foundPath)
					} else {
						fmt.Printf("%s: not found\n", extractBuiltInCommand)
					}
				}
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}
}

func findFilePath(command string) string {
	pathEnv := os.Getenv("PATH")
	directories := filepath.SplitList(pathEnv)
	for _, directory := range directories {
		fullPath := filepath.Join(directory, command)
		_, err := os.Stat(fullPath)
		if err == nil {
			return fullPath

		}
	}
	return ""
}

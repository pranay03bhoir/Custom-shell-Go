package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	// This uses the bufio package to read the user input from the command line interface.
	scanner := bufio.NewScanner(os.Stdin)
	// Implemented do while loop using for loop so that the application keeps on running until user enter "exit" to exit out of the shell.
	for {
		// This line prints $ on every line just like any other shell.
		fmt.Print("$ ")
		// Checks if user the entered any command in the shell. the scanner.Scan() returns a boolean value.
		if scanner.Scan() {
			// Extracts the text entered by the user.
			command := scanner.Text()
			// Trims the text of any trailing spaces.
			command = strings.TrimSpace(command)
			// The code exits and the loop stops when the user enter the "exit" command
			if command == "exit" {
				return
			} else if strings.HasPrefix(command, "echo ") { // this code is a implementation fo the echo command.
				fmt.Println(command[5:]) // this line prints the string entered after the echo command. example - echo hello --> hello
			} else if strings.HasPrefix(command, "type ") {
				// This code implements the type command which return if a given command is a builtin command.
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
					// This code checks whether the command has a executable path or not.
					foundPath := findFilePath(extractBuiltInCommand)
					if foundPath != "" {
						fmt.Printf("%s is %s\n", extractBuiltInCommand, foundPath)
					} else {
						fmt.Printf("%s: not found\n", extractBuiltInCommand)
					}
				}
			} else if command != "" {
				parts := strings.Fields(scanner.Text())
				command = parts[0]
				args := parts[1:]
				cmd := exec.Command(command, args...)
				args = append(args, command)
				fmt.Println("Program was passed", len(args), "args (including program name).")
				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				fmt.Println(string(output))
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}
}

// This is the implementation fo the code where the function is getting the PATH from the OS and checking if the command provided has a executable file path.
func findFilePath(command string) string {
	pathEnv := os.Getenv("PATH")
	directories := filepath.SplitList(pathEnv)
	for _, directory := range directories {
		fullPath := filepath.Join(directory, command)
		info, err := os.Stat(fullPath)
		if err == nil {
			isItRegularFile := info.Mode().IsRegular()
			isItExecutable := info.Mode().Perm()&0111 != 0
			if isItRegularFile && isItExecutable {
				return fullPath
			}

		}
	}
	return ""
}

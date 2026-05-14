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
		// Read one line of input; EOF ends the shell.
		if !scanner.Scan() {
			return
		}
		command := strings.TrimSpace(scanner.Text())

		switch {
		// The code exits and the loop stops when the user enter the "exit" command
		case command == "exit":
			return
		case strings.HasPrefix(command, "echo "): // this code is a implementation fo the echo command.
			if strings.Contains(command[5:], "'") {
				cleanStr := strings.Trim(command[5:], "'")
				fmt.Println(cleanStr)
			} else {
				extractedWords := strings.Fields(command[5:])
				// newStr := strings.Join(cleanStr, "")
				cleanStr := strings.Join(extractedWords, " ")
				fmt.Printf("%s \n", cleanStr) // this line prints the string entered after the echo command. example - echo hello --> hello

			}
		case strings.HasPrefix(command, "type "):
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
			case "pwd":
				fmt.Println(extractBuiltInCommand, "is a shell builtin")
			case "cd":
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
		case command == "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error in finding current directory: ", err)
			} else {
				fmt.Println(pwd)
			}
		case strings.HasPrefix(command, "cd "):
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", command[3:])
			}
			if command[3:] == "~" || command[3:] == "$HOME" {
				err = os.Chdir(homeDir)
				if err != nil {
					fmt.Printf("cd: %s: No such file or directory\n", command[3:])
				}
			} else {
				err = os.Chdir(command[3:])
				if err != nil {
					fmt.Printf("cd: %s: No such file or directory\n", command[3:])
				}
			}

		case command != "": // This code block executes external commands.

			// This code returns the user input (commands) in an array from.
			// Example - java -version --> ["java","-version"]
			parts := strings.Fields(command)
			// Here we separate the actual commands and arguments of the commands. and store the command in the command variable.
			// Example - java -version --> java
			name := parts[0]
			// Here the arguments are separated and are stored in the args variable.
			args := parts[1:]
			// Here we import the os/exec module to execute external commands.
			cmd := exec.Command(name, args...)
			// This returns us the output and the error combined.
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s: command not found\n", name)
				continue
			}
			// Actual output of the command is displayed.
			fmt.Print(string(output))
		default:
			// empty line: do nothing
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

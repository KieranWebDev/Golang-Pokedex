package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)

		if len(cleanedText) == 0 {
			continue
		}

		userInput := cleanedText[0]
		possibleCommands := getCommands()
		command, ok := possibleCommands[userInput]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		command.callback()
		// switch command {
		// case "help":
		// 	fmt.Println("this is the help menu")
		// case "exit":
		// 	os.Exit(0)
		// default:
		// 	fmt.Println("invalid command")
		// }

		// fmt.Println("echo", command)

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}

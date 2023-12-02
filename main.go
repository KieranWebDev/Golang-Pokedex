package main

import (
    // "bufio"
    // "fmt"
    // "os"
	// "strings"
)

// type cliCommand struct {
// 	name        string
// 	description string
// 	callback    func() error
// }

// var commands = map[string]cliCommand{
// 	"help": {
//         name:        "help",
//         description: "Displays a help message",
//         callback:    commandHelp,
//     },
//     "exit": {
//         name:        "exit",
//         description: "Exit the Pokedex",
//         callback:    commandExit,
//     },
// }

func main() {
//     scanner := bufio.NewScanner(os.Stdin)
//     scanner.Scan()
//     text := scanner.Text()
// fmt.Println(text)
startRepl()

    // for {
    //     fmt.Print("Pokedex > ")
    //     text, _ := reader.ReadString('\n')
    //     input := strings.TrimSpace(text)
        
    //     cmd, ok := commands[input]
    //     if !ok {
    //         fmt.Println("Unknown command")
    //         continue
    //     }
        
    //    cmd.callback()
		
    //     if input == "exit" {
    //         break
    //     }
    // }
}

// func commandHelp() error {
//     fmt.Println(`
//         Welcome to the Pokedex!
//         Usage:
        
//         help: Displays a help message
//         exit: Exits the program
//     `)
//     return nil
// }

// func commandExit() error {
//     fmt.Println("Exiting")
//     return nil 
// }
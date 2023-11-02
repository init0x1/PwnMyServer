package main

import (
	"fmt"

	"github.com/0xAbdoAli/PwnMyServer/api"
)

func main() {
	for {
		fmt.Print("Enter a command to be executed or exit to terminate : ")
		var command string
		_, err := fmt.Scanln(&command)
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}
		if command == "exit" {
			break
		}

		apiRes, commandError, err := api.Eval(command)
		if err != nil {
			fmt.Printf("API request failed: %v\n", err)
			return
		}
		if commandError != nil {

			fmt.Println(commandError.Error)

		} else {
			fmt.Println("Command result:", apiRes.CommandResult)
		}
	}
}

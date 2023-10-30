package main

import (
	"fmt"

	"github.com/0xAbdoAli/PwnMyServer/api"
)

func main() {
	fmt.Print("Enter a command: ")
	var command string
	_, err := fmt.Scanln(&command)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	apiRes, err := api.Eval(command)
	if err != nil {
		fmt.Printf("API request failed: %v\n", err)
		return
	}

	fmt.Println(apiRes.CommandResult)
}

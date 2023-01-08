package cmd

import (
	"fmt"
)

func init() {
	Commands = append(Commands, Command{
		Run:         "help",
		Description: "show list of commands available",
		Example:     "go run . help",
		Author:      "Mohsen Majidi",
	})
}

func CmdHelpHandler() {

	fmt.Printf("\n| %-25v | %-30v | %s\n", "Command", "Run", "Description")
	fmt.Printf("----------------------------------------------------------------------------------------------------------------------------------\n")
	for _, v := range Commands {
		fmt.Printf("| %-25v | %-30v | %s\n", v.Run, v.Example, v.Description)

	}
}

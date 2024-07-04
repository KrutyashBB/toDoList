package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "Show all commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Usage :-
$ todolist add "todo item"  # Add new task
$ todolist ls               # Show remaining tasks
$ todolist del NUMBER       # Delete a task
$ todolist done NUMBER      # Complete a task
$ todolist clear            # clear list
$ todolist commands         # Show all commands
$ todolist report           # Statistics`)
	},
}

func init() {
	rootCmd.AddCommand(commandsCmd)
}

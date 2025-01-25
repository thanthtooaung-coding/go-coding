package cmd

import (
	"fmt"
	"go-todo-cli/todo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a task description")
			return
		}
		todo.AddTask(args[0])
		fmt.Println("Task added:", args[0])
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

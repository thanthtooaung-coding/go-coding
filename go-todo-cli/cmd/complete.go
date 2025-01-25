package cmd

import (
	"fmt"
	"strconv"

	"go-todo-cli/todo"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the task number to mark as done")
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task number")
			return
		}
		todo.CompleteTask(index - 1)
		fmt.Println("Task marked as done!")
	},
}

func init() {
	RootCmd.AddCommand(completeCmd)
}

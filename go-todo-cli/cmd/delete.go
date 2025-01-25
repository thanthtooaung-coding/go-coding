package cmd

import (
	"fmt"
	"strconv"

	"go-todo-cli/todo"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the task number to delete")
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task number")
			return
		}
		todo.DeleteTask(index - 1)
		fmt.Println("Task deleted successfully!")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}

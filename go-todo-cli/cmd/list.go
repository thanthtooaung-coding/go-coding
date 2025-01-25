package cmd

import (
	"fmt"
	"go-todo-cli/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := todo.ListTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for i, task := range tasks {
			status := "❌"
			if task.Done {
				status = "✅"
			}
			fmt.Printf("%d. %s %s\n", i+1, status, task.Description)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

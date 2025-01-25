package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI to-do app",
	Long:  "A simple CLI to-do application built with Go and Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Go To-Do CLI App! Use 'todo --help' for more commands.")
	},
}

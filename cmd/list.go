package cmd

import (
	"fmt"
	"os"

	"github.com/sliceking/task-manager-cli/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("something went wrong", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("you have no tasks, take a break!")
			return
		}

		fmt.Println("You have the following tasks:")
		for i, t := range tasks {
			fmt.Printf("%d: %s\n", i+1, t.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

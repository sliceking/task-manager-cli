package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sliceking/task-manager-cli/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("somethign went wront", err)
			os.Exit(1)
		}
		fmt.Printf("added \"%s\" to you task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

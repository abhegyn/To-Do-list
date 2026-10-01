package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to To-Do List",
	Long:  `Adds a new task to To-Do List.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile(
			"tasks.csv",
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
			0644,
		)
		if err != nil {
			log.Fatalln("can't add another task due to:", err)
		}
		defer f.Close()

		var task_name string
		fmt.Println("Enter the tasks name")
		fmt.Scanln(&task_name)
		w := csv.NewWriter(f)
		defer w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

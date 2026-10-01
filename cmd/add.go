/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
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
	Short: "Adds a new tasks to the to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open("tasks.csv")
		if err != nil {
			log.Fatalln("can't add another task due to:", err)
		}
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

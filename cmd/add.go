package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	taskName string
	status   string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to To-Do List",
	Long:  `Adds a new task to To-Do List.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile(
			"tasks.csv",
			os.O_WRONLY|os.O_APPEND|os.O_CREATE, //file opened to write at the end
			0644,
		)
		if err != nil {
			log.Fatalln("can't add another task due to:", err)
		}
		defer f.Close()

		// create a flag for status with default value pending
		Count++
		newId := strconv.Itoa(Count)
		tname, err := cmd.Flags().GetString("taskname") // accepts the flag var name
		stat, err := cmd.Flags().GetString("status")
		var oStat string
		if stat == "true" || stat == "True" {
			oStat = "completed"
		} else {
			oStat = "pending"
		}

		writer := csv.NewWriter(f)

		entry := []string{
			newId,
			tname,
			oStat,
		}

		// Write the row
		if err := writer.Write(entry); err != nil {
			log.Fatalln("could not write task: %w", err)
		}

		// Force the data to be written to the file
		writer.Flush()

		// Check for errors that happened during Flush
		if err := writer.Error(); err != nil {
			log.Fatalln("could not save task: %w", err)
		}

		fmt.Println("Task added successfully")
	},
}

func init() {

	addCmd.Flags().StringVarP(
		&taskName,
		"taskname",
		"t", //supposed to be one char only
		"",
		"specify the name of the task before adding it",
	)

	addCmd.Flags().StringVarP(
		&status,
		"status",
		"s", // supposed to be one char only
		"false",
		"specify whether the task is pending: false, or complete: true",
	)

	rootCmd.AddCommand(addCmd)
}

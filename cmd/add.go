package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	taskName string
	status   string
)

func getNextID(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		// If the file does not exist, the first ID is 1.
		if os.IsNotExist(err) {
			return 1, nil
		}

		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	largestID := 0
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}

		// Skip empty or incomplete rows.
		if len(record) == 0 {
			continue
		}

		// The first row is the header: ID,Task Name,Status.
		id, err := strconv.Atoi(record[0])
		if err != nil {
			// This skips the header row.
			continue
		}

		if id > largestID {
			largestID = id
		}
	}

	return largestID + 1, nil
}

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

		tname, err := cmd.Flags().GetString("taskname") // accepts the flag var name
		stat, err := cmd.Flags().GetString("status")
		var oStat string
		if stat == "true" || stat == "True" {
			oStat = "completed"
		} else {
			oStat = "pending"
		}

		writer := csv.NewWriter(f)

		nextID, err := getNextID("tasks.csv")
		if err != nil {
			log.Fatal("could not get next ID:", err)
		}

		entry := []string{
			strconv.Itoa(nextID),
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

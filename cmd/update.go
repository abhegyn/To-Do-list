package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	newStat bool
	tId     int
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Change the status of a task",

	RunE: func(cmd *cobra.Command, args []string) error {

		status := "pending"
		if newStat {
			status = "completed"
		}

		// open the file for reading.
		file, err := os.Open("tasks.csv")
		if err != nil {
			return err
		}
		defer file.Close()

		reader := csv.NewReader(file)
		reader.Comma = '\t'

		var records [][]string

		// read and preserve the header row.
		header, err := reader.Read()
		if err != nil {
			return err
		}
		records = append(records, header)

		found := false

		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			id, _ := strconv.Atoi(record[0])

			if id == tId {
				record[2] = status
				found = true
			}

			records = append(records, record)
		}

		if !found {
			return fmt.Errorf("task with ID %d was not found", tId)
		}

		// close the file before recreating it.
		file.Close()

		// recreate tasks.csv.
		file, err = os.Create("tasks.csv")
		if err != nil {
			return err
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		writer.Comma = '\t'

		for _, record := range records {
			writer.Write(record)
		}
		writer.Flush()

		if err := writer.Error(); err != nil {
			return err
		}

		fmt.Printf("Task %d updated to %s\n", tId, status)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().BoolVarP(
		&newStat,
		"newstatus",
		"n",
		false,
		"set the task to completed; omit this flag to set it to pending",
	)

	updateCmd.Flags().IntVarP(
		&tId,
		"taskid",
		"i",
		0,
		"ID of the task to update",
	)

	if err := updateCmd.MarkFlagRequired("taskid"); err != nil {
		panic(err)
	}
}

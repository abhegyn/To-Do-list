package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/abhegyn/To-Do-list/cmd"
)

func newBegin(file *os.File) {
	w := csv.NewWriter(file)
	defer w.Flush()

	fieldName := [][]string{
		{"ID", "Task Name", "Status"},
		{"1", "Remove this example task", "Pending"},
	}

	for _, field := range fieldName {
		if err := w.Write(field); err != nil {
			log.Fatalln("error writing to the file", err)
		}
	}
}

func main() {
	cmd.Execute()
	f, err := os.OpenFile(
		"tasks.csv",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644,
	)

	if err != nil {
		os.Create("tasks.csv")

		f1, _ := os.OpenFile(
			"tasks.csv",
			os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
			0644,
		)
		defer f1.Close()
		// starting the csv by field names
		newBegin(f1)
	}
	defer f.Close()

	// starting the csv by field names
	newBegin(f)
}

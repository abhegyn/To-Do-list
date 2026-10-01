package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/abhegyn/To-Do-list/cmd"
)

func new_begin(file *os.File) {
	w := csv.NewWriter(file)
	defer w.Flush()

	field_name := []string{"ID", "Task Name", "Creation", "Status"}

	if err := w.Write(field_name); err != nil {
		log.Fatalln("error writing to the file", err)
	}
}

func read_entries(file *os.File) {

}

func write_entry(file *os.File) {

}

func main() {
	cmd.Execute()
	f, err := os.Open("tasks.csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
		os.Create("tasks.csv")

		f1, _ := os.Open("tasks.csv")
		defer f1.Close()

		new_begin(f1)
	}

	defer f.Close()

	// starting the csv by field names

	new_begin(f)

}

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

	header := []string{
		"ID",
		"Task Name",
		"Status",
	}

	if err := w.Write(header); err != nil {
		log.Fatalln("error writing header:", err)
	}
}

func main() {
	_, err := os.Stat("tasks.csv")

	if os.IsNotExist(err) {
		f, err := os.OpenFile(
			"tasks.csv",
			os.O_CREATE|os.O_WRONLY,
			0644,
		)
		if err != nil {
			log.Fatal(err)
		}

		newBegin(f)
		f.Close()
	} else if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}

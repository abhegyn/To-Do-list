package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display the entire to do list",
	Run: func(cmd *cobra.Command, args []string) {

		file, err := os.Open("tasks.csv")
		if err != nil {
			log.Fatalln("can't add another task due to:", err)
		}
		defer file.Close()

		writer := tabwriter.NewWriter(
			os.Stdout, // output destination
			0,         // minimum cell width
			4,         // tab width
			2,         // padding between columns
			' ',       // padding character
			0,         // formatting flags
		)
		defer writer.Flush()

		reader := csv.NewReader(file)
		reader.Comma = '\t'

		for {
			record, err := reader.Read() // takes a singular record/row

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Some error occurred", err)
			}

			// print each field in the row.
			for i, field := range record {
				// add a tab before every field except the first.
				if i > 0 {
					fmt.Fprint(writer, "\t")
				}

				// print the field itself.
				fmt.Fprint(writer, field)
			}

			// finish the row with a newline.
			fmt.Fprintln(writer)
		}

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

}

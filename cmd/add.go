package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add task in list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := excelize.OpenFile("tasks.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		cols, _ := file.Cols("tasks")
		if cols.Next() {
			rows, _ := cols.Rows()
			indNewRow := len(rows) + 1
			file.SetCellValue("tasks", "A"+strconv.Itoa(indNewRow), strings.Join(args, " "))
			file.SetCellValue("tasks", "B"+strconv.Itoa(indNewRow), false)
			fmt.Println("Add task success!")
		}

		if err := file.SaveAs("tasks.xlsx"); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

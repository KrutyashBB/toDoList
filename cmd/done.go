package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"strconv"
)

var doneCmd = &cobra.Command{
	Use:   "done [NUMBER]",
	Short: "Complete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		numTask, errConvert := strconv.Atoi(args[0])
		numTask += 1
		if errConvert != nil {
			fmt.Println(errConvert)
			return
		}

		file, errOpenFile := excelize.OpenFile("tasks.xlsx")
		if errOpenFile != nil {
			fmt.Println(errOpenFile)
			return
		}
		defer file.Close()

		cols, _ := file.Cols("tasks")
		if cols.Next() {
			rows, _ := cols.Rows()
			countRows := len(rows)
			if numTask > 1 && numTask <= countRows {
				file.SetCellValue("tasks", "B"+strconv.Itoa(numTask), true)
				fmt.Println("Marked task #" + strconv.Itoa(numTask-1) + " as done.")
			} else {
				fmt.Println("Wrong number task")
			}
		}

		if err := file.SaveAs("tasks.xlsx"); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

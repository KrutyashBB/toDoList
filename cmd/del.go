package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"strconv"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del [NUMBER]",
	Short: "Delete task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		numTask, errConvert := strconv.Atoi(args[0])
		numTask += 1 //Skip title
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
		countRows := 0
		if cols.Next() {
			rows, _ := cols.Rows()
			countRows = len(rows)
		}
		if numTask > 1 && numTask <= countRows {
			err := file.RemoveRow("tasks", numTask)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Deleted task #", numTask-1)
		} else {
			fmt.Println("Wrong number")
		}

		if err := file.SaveAs("tasks.xlsx"); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}

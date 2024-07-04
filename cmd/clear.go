package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all data",
	Run: func(cmd *cobra.Command, args []string) {
		file, errOpenFile := excelize.OpenFile("tasks.xlsx")
		if errOpenFile != nil {
			fmt.Println(errOpenFile)
			return
		}
		defer file.Close()

		cols, _ := file.Cols("tasks")
		countTasks := 0
		if cols.Next() {
			rows, _ := cols.Rows()
			countTasks = len(rows) - 1
		}
		for i := 0; i < countTasks; i++ {
			err := file.RemoveRow("tasks", 2)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		fmt.Println("clear success")

		if err := file.SaveAs("tasks.xlsx"); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

}

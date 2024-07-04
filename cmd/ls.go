package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"strconv"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show remaining tasks",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := excelize.OpenFile("tasks.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		indRow := 0
		rows, _ := file.Rows("tasks")
		for rows.Next() {
			if indRow == 0 { //Skip title
				indRow += 1
				continue
			}
			cols, _ := rows.Columns()
			fmt.Print("["+strconv.Itoa(indRow)+"]", cols[0])
			if cols[1] == "FALSE" {
				color.Yellow(" ...in progress")
			} else {
				color.Green(" Done!")
			}
			indRow += 1
		}
		if indRow == 1 {
			fmt.Println("No tasks found")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

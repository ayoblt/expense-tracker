/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ayoblt/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show the total amount spent",
	Long: `Calculate the sum of all expenses currently in the database.
This command does the math and tells you exactly how much money has left your pocket.`,
	Run: func(cmd *cobra.Command, args []string) {
		storage := storage.NewStorage(DBFile)
		sum, err := storage.Summary()
		if err != nil {
			fmt.Printf("An error occured: %v", err.Error())
		}

		fmt.Printf("Total expenses: ₦%d\n", sum)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ayoblt/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

var description string
var amount int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Record a new expense",
	Long: `Add a new expense to your tracker. 
You need to provide a description and the amount (in Naira).

Example:
  expense-tracker add --description "Lunch at Chicken Republic" --amount 3500`,
	Run: func(cmd *cobra.Command, args []string) {
		stg := storage.NewStorage(DBFile)

		if description == "" || amount == 0 {
			fmt.Println("Error: You must provide a --description and an --amount")
			return
		}

		expenseID, err := stg.Save(description, amount)
		if err != nil {
			fmt.Printf("Could not save expense: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Expense added successfully (ID: %d)\n", expenseID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description for this expense")

	addCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Expense amount")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

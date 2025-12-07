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

var id int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove an expense by ID",
	Long: `Delete an expense from your history using its ID.
Use the 'list' command first to find the ID of the item you want to remove.

Example:
  expense-tracker delete --id 4`,
	Run: func(cmd *cobra.Command, args []string) {
		stg := storage.NewStorage(DBFile)
		err := stg.Delete(id)
		if err != nil {
			fmt.Printf("Could not delete expense: %v", err)
			os.Exit(1)
		}

		fmt.Println("Expense deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVar(&id, "id", 0, "expense ID")
	deleteCmd.MarkFlagRequired("id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

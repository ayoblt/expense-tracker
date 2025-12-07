/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ayoblt/expense-tracker/internal/storage"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your expenses",
	Long: `Display a clean table of all your recorded expenses. 
It shows the ID, what you bought, and how much it cost. 

Great for reviewing your life choices before the month ends.`,
	Run: func(cmd *cobra.Command, args []string) {
		stg := storage.NewStorage(DBFile)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		expenses, err := stg.List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Fprintln(w, "ID\tDescription\tAmount\t")
		for _, expense := range expenses {

			fmt.Fprintf(w, "%d\t%s\t₦%d\t\n", expense.ID, expense.Description, expense.Amount)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

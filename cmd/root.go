package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "Simple expense tracker for finance needs",
	Long:  "Expense Tracker is CLI application which helps people manage their finances easily and in fast way",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(AddExpense, ListExpenses, PrintSummary, DeleteExpense, ExportCSV, UpdateExpense)
	AddExpense.Flags().StringVar(&Description, "description", "", "Flag for including description for expense")
	AddExpense.Flags().Float64Var(&Amount, "amount", 0, "Flag for including amount of the expense")
	AddExpense.Flags().StringVar(&ExpCategory, "category", "", "Flag for including category of expense")
	PrintSummary.Flags().Uint8Var(&Month, "month", 0, "Flag for including needed month statistics")
	ListExpenses.Flags().StringVar(&ExpCategory, "category", "", "Flag for including category of expense")
	DeleteExpense.Flags().IntVar(&Id, "id", 0, "Flag for identifying expense by id(to delete)")
	UpdateExpense.Flags().IntVar(&Id, "id", 0, "Flag for identifying expense by id(to update)")
	UpdateExpense.Flags().Float64Var(&Amount, "amount", 0, "Flag for updating amount of the expense")
	UpdateExpense.Flags().StringVar(&Description, "description", "", "Flag for including description for expense (to update)")
	AddExpense.MarkFlagRequired("amount")
	UpdateExpense.MarkFlagRequired("id")
	UpdateExpense.MarkFlagsOneRequired("amount", "description")
}

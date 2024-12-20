package internal

import (
	"fmt"
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cmd

import (
	"expense_tracker/funcs"
	"expense_tracker/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Description string
	Amount      float64
	Month       uint8
	Id          int
	ExpCategory string
)

var AddExpense = &cobra.Command{
	Use:   "add",
	Short: "Command for adding expenses",

	RunE: func(cmd *cobra.Command, args []string) error {

		_, err := funcs.OpenAndUnboxFile()
		if err != nil {
			return err
		}

		if Amount < 0 {
			return fmt.Errorf("amount couldn't be a negative number")
		}

		//creating expense
		newExpense := internal.NewExpense(Description, Amount, ExpCategory)

		//adding expense to the list
		internal.ListOfExpenses = append(internal.ListOfExpenses, newExpense)

		if err = funcs.SaveExpenses(); err != nil {
			return err
		}

		fmt.Printf("Expense added successfully (ID: %d)\n", newExpense.ID)
		return nil
	},
}

var ListExpenses = &cobra.Command{
	Use:   "list",
	Short: "List is a command for listing all expenses",

	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := funcs.OpenAndUnboxFile()
		if err != nil {
			return err
		}
		defer file.Close()

		internal.ListOfExpenses.PrintExpenses(ExpCategory)
		return nil
	},
}

var PrintSummary = &cobra.Command{
	Use:   "summary",
	Short: "Command for summarizing expenses",

	RunE: func(cmd *cobra.Command, args []string) error {

		file, err := funcs.OpenAndUnboxFile()
		if err != nil {
			return err
		}
		defer file.Close()

		internal.ListOfExpenses.ExpenseSummary(Month)

		return nil
	},
}

var DeleteExpense = &cobra.Command{
	Use:   "delete",
	Short: "Command for deleting expenses from expense list",

	RunE: func(cmd *cobra.Command, args []string) error {
		//open file, read it and write its contents to slice
		_, err := funcs.OpenAndUnboxFile()
		if err != nil {
			return err
		}

		//deleting an element
		err = internal.ListOfExpenses.DeleteExpenseFunc(Id)
		if err != nil {
			return err
		}

		//saving slice to file
		if err = funcs.SaveExpenses(); err != nil {
			return err
		}
		return nil
	},
}

var ExportCSV = &cobra.Command{
	Use:   "export",
	Short: "Command for exporting expenses list as a CSV file",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {

		file, err := funcs.OpenAndUnboxFile()
		if err != nil {
			return err
		}
		defer file.Close()

		//check if list with expenses is empty
		if internal.ListOfExpenses.IsEmpty() {
			return fmt.Errorf("expense list is empty")
		}

		filename := args[0]

		//creating csv file or truncating file if was already created
		textFile, err := funcs.CreateCSV(filename)
		if err != nil {
			return err
		}
		defer textFile.Close()

		//saving to CSV file
		if err := funcs.SaveToCSV(textFile, internal.ListOfExpenses); err != nil {
			return err
		}
		return nil
	},
}

var UpdateExpense = &cobra.Command{
	Use:   "update",
	Short: "Command for updating expense by id",

	RunE: func(cmd *cobra.Command, args []string) error {

		_, err := funcs.OpenAndUnboxFile()
		if err != nil {
			return err
		}

		internal.ListOfExpenses.UpdateExpense(Id, Amount, Description)

		//saving slice to file
		if err = funcs.SaveExpenses(); err != nil {
			return err
		}
		return nil
	},
}

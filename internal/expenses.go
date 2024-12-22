package internal

import (
	"fmt"
	"slices"
	"time"
)

type ExpenseList []Expense

// array to save expenses with type
var ListOfExpenses ExpenseList

type Expense struct {
	ID       int     `json:"id"`
	Date     string  `json:"date"`
	Descr    string  `json:"description"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}

func NewExpense(descr string, amount float64, category string) Expense {
	// time formatting
	time := time.Now().Format("2006-01-02")

	// finding id of new element
	var id int
	if len(ListOfExpenses) == 0 {
		id = 1
	} else {
		id = ListOfExpenses[len(ListOfExpenses)-1].ID + 1
	}

	var newExpense = Expense{
		ID:       id,
		Date:     time,
		Descr:    descr,
		Amount:   amount,
		Category: category,
	}
	return newExpense
}

// methods for slice
func (list ExpenseList) IsEmpty() bool {
	if len(list) == 0 {
		fmt.Println("Expense list is empty")
		return true
	}
	return false
}

func (list ExpenseList) PrintExpenses(category string) {

	if list.IsEmpty() {
		return
	}

	categoriesList := list.UpdateCategories()

	if category == "" {
		fmt.Printf("ID \t %10s \t %15s \t Amount \t %10s\n", "Date", "Description", "Category")
		for _, val := range list {
			fmt.Printf("%d \t %10s \t %15s \t $%g \t\t %10s\n", val.ID, val.Date, val.Descr, val.Amount, val.Category)
		}
	} else if slices.Contains(categoriesList, category) {
		fmt.Printf("ID \t %10s \t %15s \t Amount \t %10s\n", "Date", "Description", "Category")
		for _, val := range list {
			if val.Category == category {
				fmt.Printf("%d \t %10s \t %15s \t $%g \t\t %10s\n", val.ID, val.Date, val.Descr, val.Amount, val.Category)
			}
		}
	} else {
		fmt.Printf("Category %s not found\n", category)
	}
}

func (list ExpenseList) ExpenseSummary(monthFlag uint8) {
	if list.IsEmpty() {
		return
	}

	var totalAmount float64

	if monthFlag != 0 {
		for _, expense := range list {
			parsedTime, _ := time.Parse("2006-01-02", expense.Date)

			if parsedTime.Month() == time.Month(monthFlag) {
				totalAmount += expense.Amount
			}
		}
		fmt.Printf("Total expenses for %s: $%g", time.Month(monthFlag), totalAmount)
	} else {
		for _, expense := range list {
			totalAmount += expense.Amount
		}
		fmt.Printf("Total expenses: $%g\n", totalAmount)
	}
}

func (list *ExpenseList) DeleteExpenseFunc(id int) error {

	for idx, expense := range *list {
		if expense.ID == id {
			*list = append((*list)[:idx], (*list)[idx+1:]...)
			fmt.Println("Expense deleted successfully")
			return nil
		}
	}
	return fmt.Errorf("expense with ID %d not found", id)
}

func (list ExpenseList) UpdateCategories() (categoryList []string) {
	for _, expense := range list {
		categoryList = append(categoryList, expense.Category)
	}

	return
}

func (list *ExpenseList) UpdateExpense(id int, amount float64, descr string) {
	for idx, expense := range *list {
		if expense.ID == id {
			successMessage := func() {
				fmt.Println("Expense was successfully updated")
			}
			switch {
			case amount == 0:
				(*list)[idx].Descr = descr
				successMessage()
				return

			case descr == "":
				(*list)[idx].Amount = amount
				successMessage()
				return

			default:
				(*list)[idx].Descr = descr
				(*list)[idx].Amount = amount
				successMessage()
				return

			}
		}
	}

	fmt.Printf("Expense with id %d not found\n", id)
}

# Expense Tracker - CLI Application

## Installation

```bash
git clone https://github.com/PureTeamLead/ExpenseTracker_Go.git
cd main
```

Building the app

```bash
go build -o expense-tracker
```

## What App can do?

### Add Expense

Amount flag is required here

```bash
./expense-tracker add --description Lunch --amount 20 --category Restaurants
#Expense added successfully (ID: 1)
```

### Update Expense

ID flag is required here, and one of group (description, amount) is required too.

```bash
./expense-tracker update --id 1 --description Breakfast --amount 10
#Expense was successfully updated
```

### List Expenses

```bash
./expense-tracker list
# ID	Date			Description		Amount	Category
# 1	2024-12-21	Breakfast		$10		Restaurants
```

Program can list expenses by filtering category

```bash
./expense-tracker list --category Restaurants
```

### Delete Expense

ID flag should be provided here

```bash
./expense-tracker delete --id 1
```

### Summarize Expenses

```bash
./expense-tracker summary
```

Also application can summarize expenses by provided month

```bash
./expense-tracker summary --month 12
```

package funcs

import (
	"encoding/json"
	"expense_tracker/internal"
	"fmt"
	"os"
	"path"
)

var DbFileName = "expenses.json"

func Filepath(filename string) (string, string) {
	var dirName = "data"

	cwd, _ := os.Getwd()

	dirpath := path.Join(cwd, dirName)
	filepath := path.Join(dirpath, filename)

	return dirpath, filepath
}

func OpenExpensesFile() (*os.File, error) {
	//get a filepath
	dirpath, filepath := Filepath(DbFileName)

	//check if the directory and file exist
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating a file...")

		//creating a directory
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("error creating directory")
		}

		//creating a file
		newFile, err := os.Create(filepath)
		if err != nil {
			return nil, fmt.Errorf("error creating file")
		}

		//write an empty slice to file for future includes
		newFile.Write([]byte("[]"))

		return newFile, nil
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening file")
	}
	return file, nil
}

func UnboxFile(file *os.File, list *internal.ExpenseList) error {
	fileStats, err := file.Stat()
	if err != nil {
		return err
	}

	if fileStats.Size() <= 2 {
		return nil
	}

	err = json.NewDecoder(file).Decode(list)
	if err != nil {
		return err
	}
	return nil
}

func OpenAndUnboxFile() (*os.File, error) {
	file, err := OpenExpensesFile()
	if err != nil {
		return nil, err
	}

	if err = UnboxFile(file, &internal.ListOfExpenses); err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}

func SaveExpenses() error {

	_, filepath := Filepath(DbFileName)
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(&internal.ListOfExpenses); err != nil {
		return err
	}

	return nil
}

func CreateCSV(filename string) (*os.File, error) {
	dirpath, filepath := Filepath(filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("Creating CSV file...")

		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("error creating directory for file")
		}

		newFile, err := os.Create(filepath)
		if err != nil {
			return nil, fmt.Errorf("error creating file")
		}
		return newFile, nil
	}

	file, err := os.OpenFile(filepath, os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening file")
	}
	return file, nil
}

func SaveToCSV(file *os.File, list internal.ExpenseList) error {

	for _, expense := range list {
		s := fmt.Sprintf("ID: %d, Description: %s, Date: %s, Amount: %g, Category: %s\n", expense.ID, expense.Descr, expense.Date, expense.Amount, expense.Category)
		if _, err := file.WriteString(s); err != nil {
			return err
		}
	}

	return nil
}

package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	expenseMod "github.com/ayoblt/expense-tracker/internal/expense"
)

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	return &Storage{
		filename: filename,
	}
}

func (s *Storage) Save(description string, amount int) (int, error) {
	dat, err := os.ReadFile(s.filename)

	var expenseTracker expenseMod.ExpenseTracker
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			expenseTracker = expenseMod.ExpenseTracker{NextID: 1}
		} else {
			return 0, err
		}
	} else {
		_ = json.Unmarshal(dat, &expenseTracker)
	}

	expenseId := expenseTracker.Add(description, amount)

	dat, err = json.MarshalIndent(expenseTracker, "", "  ")
	if err != nil {
		return 0, err
	}

	err = os.WriteFile(s.filename, dat, 0644)
	if err != nil {
		return 0, err
	}

	return expenseId, nil
}

func (s *Storage) List() ([]expenseMod.Expense, error) {
	dat, err := os.ReadFile(s.filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []expenseMod.Expense{}, nil
		}
		return nil, err
	}

	var expenseTracker expenseMod.ExpenseTracker
	if err := json.Unmarshal(dat, &expenseTracker); err != nil {
		return nil, fmt.Errorf("file contains corrupt JSON: %w", err)
	}

	return expenseTracker.Expenses, nil
}

func (s *Storage) Summary() (int, error) {
	expenses, err := s.List()
	if err != nil {
		return 0, err
	}

	var sum int
	for _, expense := range expenses {
		sum += expense.Amount
	}

	return sum, nil
}

func (s *Storage) Delete(id int) error {
	dat, err := os.ReadFile(s.filename)

	var expenseTracker expenseMod.ExpenseTracker
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			expenseTracker = expenseMod.ExpenseTracker{NextID: 1}
		} else {
			return err
		}
	} else {
		_ = json.Unmarshal(dat, &expenseTracker)
	}

	err = expenseTracker.Delete(id)
	if err != nil {
		return err
	}

	dat, err = json.MarshalIndent(expenseTracker, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(s.filename, dat, 0644)
	if err != nil {
		return err
	}

	return nil
}

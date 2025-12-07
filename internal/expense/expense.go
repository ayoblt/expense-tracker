package expense

import "fmt"

type Expense struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type ExpenseTracker struct {
	Expenses []Expense `json:"expenses"`
	NextID   int       `json:"next_id"`
}

func NewExpense(ID int, description string, amount int) Expense {
	newExpense := Expense{
		ID:          ID,
		Description: description,
		Amount:      amount,
	}
	return newExpense
}

func NewExpenseTracker() *ExpenseTracker {
	return &ExpenseTracker{NextID: 1}
}

func (e *ExpenseTracker) Add(description string, amount int) int {
	newExpense := NewExpense(e.NextID, description, amount)

	e.Expenses = append(e.Expenses, newExpense)
	e.NextID = newExpense.ID + 1
	return newExpense.ID
}

func (e *ExpenseTracker) Delete(id int) error {
	found := false

	newExpenses := []Expense{}
	for _, expense := range e.Expenses {
		if expense.ID == id {
			found = true
		} else {
			newExpenses = append(newExpenses, expense)
		}
	}

	if !found {
		return fmt.Errorf("ID %d not found", id)
	}

	e.Expenses = newExpenses
	return nil
}

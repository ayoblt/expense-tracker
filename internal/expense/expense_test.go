package expense

import (
	"reflect"
	"slices"
	"testing"
)

func TestExpenseCreate(t *testing.T) {
	newExpense := NewExpense(1, "Write a test", 20)

	want := Expense{
		ID:          1,
		Description: "Write a test",
		Amount:      20,
	}

	if !reflect.DeepEqual(newExpense, want) {
		t.Errorf("got %v, want %+v", newExpense, want)
	}
}

func TestExpenseAdd(t *testing.T) {
	expenseTracker := NewExpenseTracker()

	id := expenseTracker.Add("Lunch", 20)

	if id != 1 {
		t.Errorf("got %d, want %d", id, 1)
	}
}

func TestExpenseDelete(t *testing.T) {
	expenseTracker := NewExpenseTracker()

	_ = expenseTracker.Add("Breakfast", 30)
	_ = expenseTracker.Add("Lunch", 20)
	_ = expenseTracker.Add("Dinner", 10)

	err := expenseTracker.Delete(2)
	if err != nil {
		t.Fatalf("unexpected error in deletion: %v", err)
	}
	got := expenseTracker.Expenses
	want := []Expense{
		{ID: 1, Description: "Breakfast", Amount: 30},
		{ID: 3, Description: "Dinner", Amount: 10},
	}

	if !slices.Equal(want, got) {
		t.Errorf("got %v want %v", got, want)
	}

}

func TestExpenseList(t *testing.T) {
	expenseTracker := NewExpenseTracker()

	_ = expenseTracker.Add("Breakfast", 10)
	got := expenseTracker.Expenses
	want := []Expense{
		{ID: 1, Description: "Breakfast", Amount: 10},
	}

	if !slices.Equal(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

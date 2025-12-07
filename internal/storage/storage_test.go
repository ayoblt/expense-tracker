package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	expenseMod "github.com/ayoblt/expense-tracker/internal/expense"
)

func TestNewStorage(t *testing.T) {
	tmpDir := t.TempDir()

	tmpFile := filepath.Join(tmpDir, "test_db.json")

	storage := NewStorage(tmpFile)

	expenseDescription := "Create storage"
	expenseAmount := 20

	_, err := storage.Save(expenseDescription, expenseAmount)
	if err != nil {
		t.Fatalf("unexpected error saving: %v", err)
	}

	var loaded expenseMod.ExpenseTracker
	dat, err := os.ReadFile(storage.filename)
	if err != nil {
		t.Fatalf("file was not created: %v", err)
	}

	if err := json.Unmarshal(dat, &loaded); err != nil {
		t.Errorf("file not a valid JSON: %v", err)
	}

	if len(loaded.Expenses) != 1 {
		t.Fatalf("expected 1 got %d", len(loaded.Expenses))
	}

	if loaded.Expenses[0].Description != expenseDescription {
		t.Errorf("description err: got %s want %s", loaded.Expenses[0].Description, expenseDescription)
	}
}

func TestStorageList(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_db.json")

	initialData := []byte(`{
		"expenses": [
			{ "id": 1, "description": "Previous Expense 1", "amount": 3000 },
			{ "id": 2, "description": "Previous Expense 2", "amount": 4000 }
	],
	"next_id": 3
	}`)

	if err := os.WriteFile(tmpFile, initialData, 0644); err != nil {
		t.Fatalf("could not seed test data: %v", err)
	}

	storage := NewStorage(tmpFile)

	expenses, err := storage.List()
	if err != nil {
		t.Fatalf("unexpected error listing error: %v", err)
	}

	if len(expenses) != 2 {
		t.Fatalf("got %d want 2", len(expenses))
	}

	if expenses[0].Description != "Previous Expense 1" {
		t.Errorf("expected first expense to be 'Previous Expense 1', got '%s'", expenses[0].Description)
	}
}

func TestStorageSummary(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_db.json")

	initialData := []byte(`{
		"expenses": [
			{ "id": 1, "description": "Previous Expense 1", "amount": 3000 },
			{ "id": 2, "description": "Previous Expense 2", "amount": 4000 }
	],
	"next_id": 3
	}`)

	if err := os.WriteFile(tmpFile, initialData, 0644); err != nil {
		t.Fatalf("could not seed test data: %v", err)
	}

	storage := NewStorage(tmpFile)

	got, err := storage.Summary()
	if err != nil {
		t.Fatalf("unexpected error with summary: %v", err)
	}

	want := 7000

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestStorageDelete(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test_db.json")

	initialData := []byte(`{
		"expenses": [
			{ "id": 1, "description": "Previous Expense 1", "amount": 3000 },
			{ "id": 2, "description": "Previous Expense 2", "amount": 4000 },
			{ "id": 3, "description": "Previous Expense 3", "amount": 2000 }
	],
	"next_id": 4
	}`)

	if err := os.WriteFile(tmpFile, initialData, 0644); err != nil {
		t.Fatalf("could not seed test data: %v", err)
	}

	storage := NewStorage(tmpFile)

	err := storage.Delete(2)
	if err != nil {
		t.Fatalf("unexpected error with deletion: %v", err)
	}

	dat, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("could not open test file: %v", err)
	}
	var expenseTracker expenseMod.ExpenseTracker
	if err := json.Unmarshal(dat, &expenseTracker); err != nil {
		t.Fatalf("JSON file corrupted: %v", err)
	}

	if len(expenseTracker.Expenses) != 2 {
		t.Errorf("got %d, want 2", len(expenseTracker.Expenses))
	}

	if expenseTracker.Expenses[0].ID == 2 || expenseTracker.Expenses[1].ID == 2 {
		t.Errorf("expense not deleted, got %v", expenseTracker.Expenses)
	}
}

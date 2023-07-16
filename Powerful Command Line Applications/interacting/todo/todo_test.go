package todo_test

import (
	"os"
	"testing"

	"github.com/hellskater/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l[0].Task)
	}

}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"

	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l[0].Task)
	}

	l.Delete(1)

	if len(l) != 0 {
		t.Errorf("Task should be deleted")
	}
}

func TestSave(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"

	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q", taskName, l1[0].Task)
	}

	tempFile, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tempFile.Name())

	if err := l1.Save(tempFile.Name()); err != nil {
		t.Fatalf("Error saving to file: %s", err)
	}

	if err := l2.Get(tempFile.Name()); err != nil {
		t.Fatalf("Error getting from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Expected %q, got %q", l1[0].Task, l2[0].Task)
	}

}

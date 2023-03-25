package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateTables(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		t.Fatal(err)
	}
}

func TestCreateCategorySummary(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err = createTables(db); err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO tasks (title, value, category) VALUES (?, ?, ?)", "Task 1", 10, "A")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO tasks (title, value, category) VALUES (?, ?, ?)", "Task 2", 20, "A")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO tasks (title, value, category) VALUES (?, ?, ?)", "Task 3", 30, "B")
	if err != nil {
		t.Fatal(err)
	}

	summaries, err := createCategorySummary(db)
	if err != nil {
		t.Fatal(err)
	}

	if len(summaries) != 2 {
		t.Fatalf("expected 2 summaries, got %d", len(summaries))
	}
	if summaries[0].Title != "A" {
		t.Errorf("expected A, got %s", summaries[0].Title)
	}
	if summaries[0].Tasks != 2 {
		t.Errorf("expected 2, got %d", summaries[0].Tasks)
	}
	if summaries[0].AvgValue != 15 {
		t.Errorf("expected 15, got %f", summaries[0].AvgValue)
	}
	if summaries[1].Title != "B" {
		t.Errorf("expected B, got %s", summaries[1].Title)
	}
	if summaries[1].Tasks != 1 {
		t.Errorf("expected 1, got %d", summaries[1].Tasks)
	}
	if summaries[1].AvgValue != 30 {
		t.Errorf("expected 30, got %f", summaries[1].AvgValue)
	}
}

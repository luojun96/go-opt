package main

import "database/sql"

type CategorySummary struct {
	Title    string
	Tasks    int
	AvgValue float64
}

func createTables(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY,title TEXT, value INTEGER, category TEXT)")
	return err
}

func createCategorySummary(db *sql.DB) ([]CategorySummary, error) {
	rows, err := db.Query("SELECT category, COUNT(*), AVG(value) FROM tasks GROUP BY category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []CategorySummary
	for rows.Next() {
		var s CategorySummary
		if err := rows.Scan(&s.Title, &s.Tasks, &s.AvgValue); err != nil {
			return nil, err
		}
		summaries = append(summaries, s)
	}
	return summaries, nil
}

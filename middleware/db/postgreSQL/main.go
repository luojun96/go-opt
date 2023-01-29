package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	// _ "github.com/jackc/pgx/v4/stdlib"

	"github.com/jackc/pgx/v4"
)

type Shoe struct {
	Number string
	Name   string
	Eur    float64
}

type RowSrc struct {
	cr     *csv.Reader
	values []interface{}
	err    error
}

func (r *RowSrc) Next() bool {
	record, err := r.cr.Read()
	if errors.Is(err, io.EOF) {
		// r.err = err
		return false
	}
	if err != nil {
		r.err = err
		return false
	}
	r.values = make([]interface{}, 3)
	r.values[0] = record[0]
	r.values[1] = record[1]
	r.values[2], _ = strconv.ParseFloat(record[2], 64)
	// r.values[2] = float64(record[2])
	return true
}

// Values returns the values for the current row.
func (r *RowSrc) Values() ([]interface{}, error) {
	return r.values, r.err
}

// Err returns any error that has been encountered by the CopyFromSource. If
// this is not nil *Conn.CopyFrom will abort the copy.
func (r *RowSrc) Err() error {
	return r.err
}

func NewConn() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "207.148.113.13:5432",
		User:   url.UserPassword("jun", "Initial0!"),
		Path:   "shoe",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect %w", err)
	}
	return conn, nil
}

func InsertShoesByPgx(conn *pgx.Conn, shoes []Shoe) error {
	if err := conn.BeginFunc(context.Background(), func(tx pgx.Tx) error {
		for _, shoe := range shoes {
			_, err := tx.Exec(context.Background(), "INSERT INTO shoes(number, name, eur) VALUES($1, $2, $3)", shoe.Number, shoe.Name, shoe.Eur)
			if err != nil {
				return fmt.Errorf("tx.ExecContext %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("conn.BeginFunc %w", err)
	}
	return nil
	// tx, err := conn.BeginTx(context.Background(), nil)
	// if err != nil {
	// 	return fmt.Errorf("conn.BeginTx %w", err)
	// }

	// defer func() {
	// 	_ = tx.Rollback(context.Background())
	// }()

	// if err = tx.Commit(context.Background()); err != nil {
	// 	return fmt.Errorf("tx.Commit %w", err)
	// }
	// return nil

}

func InsertShoesFromLocalFile(conn *pgx.Conn) error {
	f, err := os.Open(path.Join("resources", "data.csv"))
	if err != nil {
		return fmt.Errorf("os.Open %w", err)
	}

	cr := csv.NewReader(f)

	if err = conn.BeginTxFunc(context.Background(), pgx.TxOptions{}, func(tx pgx.Tx) error {
		for {
			record, err := cr.Read()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				return fmt.Errorf("cr.Read %w", err)
			}
			if _, err = conn.Exec(context.Background(), "INSERT INTO shoes(number, name, eur) VALUES($1, $2, $3)", record[0], record[1], record[2]); err != nil {
				return fmt.Errorf("conn.Exec %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("conn.BeginTxFunc %w", err)
	}
	return nil
}

func InsertShoesFromLocalFileInBatch(conn *pgx.Conn) error {
	f, err := os.Open(path.Join("resources", "data.csv"))
	if err != nil {
		return fmt.Errorf("os.Open %w", err)
	}
	cr := csv.NewReader(f)

	rowSrc := &RowSrc{
		cr: cr,
	}
	count, err := conn.CopyFrom(context.Background(), pgx.Identifier{"shoes"}, []string{"number", "name", "eur"}, rowSrc)
	if err != nil {
		return fmt.Errorf("conn.CopyFrom %w", err)
	}
	fmt.Println("Insert rows", count)
	return nil
}

func main() {
	conn, err := NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = conn.Close(context.Background())
		fmt.Println("closed")
	}()

	now := time.Now()

	if err = InsertShoesFromLocalFileInBatch(conn); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Total time: %s\n", time.Since(now).String())

	// if err = InsertShoesByPgx(conn, []Shoe{
	// 	{
	// 		Number: "04",
	// 		Name:   "AJ2",
	// 		Eur:    42.0,
	// 	},
	// 	{
	// 		Number: "05",
	// 		Name:   "AJ5 OW",
	// 		Eur:    41.0,
	// 	},
	// }); err != nil {
	// 	fmt.Println(err)
	// }
}

// func newDB() (*sql.DB, error) {
// 	dsn := url.URL{
// 		Scheme: "postgres",
// 		Host:   "207.148.113.13:5432",
// 		User:   url.UserPassword("jun", "Initial0!"),
// 		Path:   "shoe",
// 	}
// 	q := dsn.Query()
// 	q.Add("sslmode", "disable")
// 	dsn.RawQuery = q.Encode()

// 	db, err := sql.Open("pgx", dsn.String())
// 	if err != nil {
// 		return nil, fmt.Errorf("sql.Open %w", err)
// 	}

// 	return db, nil
// }

// func InsertShoes(db *sql.DB, shoes []Shoe) error {
// 	tx, err := db.BeginTx(context.Background(), nil)
// 	if err != nil {
// 		return fmt.Errorf("db.BeginTx %w", err)
// 	}
// 	defer func() {
// 		_ = tx.Rollback()
// 	}()

// 	for _, shoe := range shoes {
// 		_, err = tx.ExecContext(context.Background(), "INSERT INTO shoes(number, name, eur) VALUES($1, $2, $3)", shoe.Number, shoe.Name, shoe.Eur)
// 		if err != nil {
// 			return fmt.Errorf("tx.ExecContext %w", err)
// 		}
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("tx.Commit %w", err)
// 	}

// 	return nil
// }

// func main() {
// 	db, err := newDB()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer func() {
// 		db.Close()
// 		fmt.Println("DB Close")
// 	}()

// 	if err = InsertShoes(db, []Shoe{
// 		{
// 			Number: "02",
// 			Name:   "AJ1",
// 			Eur:    42.0,
// 		},
// 		{
// 			Number: "03",
// 			Name:   "AJ4",
// 			Eur:    42.5,
// 		},
// 	}); err != nil {
// 		fmt.Println("Insert Shoes", err)
// 	}

// }

// just sql statment
func oldMain() {
	// db, err := newDB()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// ctx := context.Background()
	// if err := db.PingContext(ctx); err != nil {
	// 	fmt.Println("db.PingContext", err)
	// 	return
	// }

	// row := db.QueryRowContext(ctx, "SELECT name FROM shoes WHERE number = '01'")
	// if err = row.Err(); err != nil {
	// 	fmt.Println("db Query", err)
	// 	return
	// }

	// var name string
	// if err = row.Scan(&name); err != nil {
	// 	fmt.Println("row.Scan", err)
	// 	return
	// }

	// fmt.Println("name", name)

	// // _, err = db.ExecContext(context.Background(), "INSERT INTO shoes(number, name, eur) VALUES('02', 'DUNK LOW', 38.5)")
	// // number, name, eur := "03", "AJ1 HIGH OG", 41
	// // _, err = db.ExecContext(context.Background(), "INSERT INTO shoes(number, name, eur) VALUES($1, $2, $3)", number, name, eur)
	// // if err != nil {
	// // 	fmt.Println("db.ExecContext", err)
	// // 	return
	// // }

	// rows, err := db.QueryContext(context.Background(), "SELECT * from shoes")
	// if err != nil {
	// 	fmt.Println("row.Scan", err)
	// 	return
	// }

	// defer func() {
	// 	_ = rows.Close()
	// }()

	// if rows.Err() != nil {
	// 	fmt.Println("row.Err", rows.Err())
	// 	return
	// }

	// for rows.Next() {
	// 	var number string
	// 	var name string
	// 	var eur float64

	// 	err = rows.Scan(&number, &name, &eur)
	// 	if err != nil {
	// 		fmt.Println("rows.Scan", err)
	// 		return
	// 	}

	// 	fmt.Println("number", number, "name", name, "eur", eur)
	// }

}

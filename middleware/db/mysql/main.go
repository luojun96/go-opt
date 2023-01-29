package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func main() {
	c := mysql.Config{
		User:          "root",
		Passwd:        "Initial0!",
		Net:           "tcp",
		Addr:          "207.148.113.13",
		DBName:        "shoes",
		AllowAllFiles: true,
	}
	fmt.Println(c.FormatDSN())

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	defer func() {
		db.Close()
		fmt.Println("closed")
	}()

	if err = db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	filePath := path.Join("resources", "data.csv")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open", err)
		return
	}
	defer func() {
		_ = db.Close()
	}()

	mysql.RegisterReaderHandler("data", func() io.Reader {
		buffer := bytes.Buffer{}
		csvWriter := csv.NewWriter(&buffer)

		csvReader := csv.NewReader(file)
		for {
			record, err := csvReader.Read()
			if errors.Is(err, io.EOF) {
				break
			}

			if err != nil {
				fmt.Println("cscReader.Read", err)
				return nil
			}

			record[1] = strings.ToLower(record[1])

			csvWriter.Write(record)
		}
		csvWriter.Flush()
		return &buffer
	})

	_, err = db.ExecContext(context.Background(),
		"LOAD DATA LOCAL INFILE 'Reader::data' INTO TABLE shoes "+
			`FIELDS TERMINATED BY ','
			ENCLOSED BY '"'
			(number, name)`)
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return
	}

	// filePath := path.Join("resources", "data.csv")
	// mysql.RegisterLocalFile(filePath)
	// _, err = db.ExecContext(context.Background(),
	// 	"LOAD DATA LOCAL INFILE '"+filePath+"' INTO TABLE shoes "+
	// 		`FIELDS TERMINATED BY ','
	// 		ENCLOSED BY '"'
	// 		(number, name)`)
	// if err != nil {
	// 	fmt.Println("db.ExecContext", err)
	// 	return
	// }
	// row := db.QueryRowContext(context.Background(), "SELECT name FROM shoes WHERE number = ?", "01")
	// if row.Err() != nil {
	// 	fmt.Println("db.QueryRowContext", row.Err())
	// 	return
	// }

	// var name string
	// if err = row.Scan(&name); err != nil {
	// 	fmt.Println("row.Scan", err)
	// 	return
	// }

	// fmt.Printf("name = %s\n", name)

}

package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4"
)

func TestInsertShoesByPgx(t *testing.T) {
}

func TestInsertShoesFromLocalFile(t *testing.T) {
	type args struct {
		conn *pgx.Conn
	}
	conn, err := NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = conn.Close(context.Background())
		fmt.Println("closed")
	}()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				conn: conn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertShoesFromLocalFile(tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("InsertShoesFromLocalFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestInsertShoesFromLocalFileInBatch(t *testing.T) {
	type args struct {
		conn *pgx.Conn
	}
	conn, err := NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		_ = conn.Close(context.Background())
		fmt.Println("closed")
	}()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				conn: conn,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertShoesFromLocalFileInBatch(tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("InsertShoesFromLocalFileInBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

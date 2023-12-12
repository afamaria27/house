package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"house/information/createShowHouse"
	"os"
)

func main() {
	createShowHouse.ShowHouse()

	urlExample := "postgres://postgres:1111@localhost:5432/test_db"
	_, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("ready")
	}
}

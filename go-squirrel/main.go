package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
)

func main() {
	connStr := "postgres://postgres:password@localhost:5433/mydb?sslmode=disable"

	// Open the PostgreSQL connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}
	fmt.Println("Connected to PostgreSQL!")

	// Create table if not exists
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(50),
        age INT
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}

	// Insert data using Squirrel
	insertBuilder := squirrel.Insert("users").
		Columns("name", "age").
		Values("Alice", 30).
		Values("Bob", 25).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err := insertBuilder.ToSql()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(sql, args...)
	if err != nil {
		log.Fatal("Failed to insert data: ", err)
	}
	fmt.Println("Data inserted successfully!")

	// Query data using Squirrel
	selectBuilder := squirrel.Select("id", "name", "age").
		From("users").
		Where(squirrel.Gt{"age": 20}).
		PlaceholderFormat(squirrel.Dollar)

	sql, args, err = selectBuilder.ToSql()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(sql, args...)
	if err != nil {
		log.Fatal("Failed to query data: ", err)
	}
	defer rows.Close()

	fmt.Println("Query Results:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

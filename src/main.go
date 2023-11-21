package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Person struct {
	id string
	name string
	email string
	createdAt sql.NullTime
	updatedAt sql.NullTime
}

func getAll() []Person {
	// Connect to the database
	//
	connStr := "user=your_user password=your_password dbname=database_name" // Change this to your database
	db, err := sql.Open("postgres", connStr)

	// Check if it connected propertly
	//
	if err != nil {
		log.Fatal(err)
		return []Person {}
	}

	// Searches for data in the database
	//
	queryString := `SELECT id, name, email, "createdAt", "updatedAt" FROM "Person"`
	rows, err := db.Query(queryString)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []Person

	// Gets rows from database
	// and maps it to a struct
	//
	for rows.Next() {
		var person Person

		if err := rows.Scan(
			&person.id,
			&person.name,
			&person.email,
			&person.createdAt,
			&person.updatedAt); err != nil {
				log.Fatal(err)
			}

		result = append(result, person)
	}

	return result
}

func main() {
	var personList []Person = getAll()
	log.Printf("%+v\n", personList)
}

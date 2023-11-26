package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Person struct {
	Id        string
	Name      string
	Email     string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func getConnection() *sql.DB {
	connStr := "user=your_user password=your_password dbname=database_name" // Change this to your database
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getAll(db *sql.DB) []Person {
	if db == nil {
		return nil
	}

	// Searches for data in the database
	//
	queryString := `SELECT id, name, email, "createdAt", "updatedAt" FROM "Person"`
	rows, err := db.Query(queryString)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var result []Person

	// Gets rows from database
	// and maps it to a struct
	//
	for rows.Next() {
		var person Person

		if err := rows.Scan(
			&person.Id,
			&person.Name,
			&person.Email,
			&person.CreatedAt,
			&person.UpdatedAt); err != nil {
				log.Fatal(err)
			}

		result = append(result, person)
	}

	return result
}

func main() {
	var connection = getConnection();

	var personList []Person = getAll(connection)
	log.Printf("%+v\n", personList)
}

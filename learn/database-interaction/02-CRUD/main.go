package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loding .env file")
	}

	connStr := os.Getenv("DATABASE_URI")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        age INT
    );`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table 'users' created successfully!")

	query = `
    INSERT INTO users (name, email, age) VALUES ('Gaurav', 'gaurav@gmail.com', 34)
  `

	data, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Data inset successfully in user table", data)

	query = `
    UPDATE users 
    SET email = 'new-gaurav@gmail.com'
    WHERE email = 'gaurav@gmail.com'
  `

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print("Data UPDATE successfully")

	query = `
    DELETE FROM users
    WHERE email = 'new-gaurav@gmail.com'
  `

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print("Data DELETE successfully")
}

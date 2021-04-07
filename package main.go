package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	var choise int
	fmt.Print("Аuthorization - 0, registration - 1: ")
	fmt.Scan(&choise)

	if choise == 0 {
		var email string
		var password string

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter your password: ")
		fmt.Scan(&password)

		counter := query(email, password)
		if counter != 0 {
			fmt.Print("you entered")
		} else {
			fmt.Print("incorrect email or password")
		}
	}

	if choise == 1 {
		var first_name string
		var last_name string
		var email string
		var password string

		fmt.Print("Enter your first name: ")
		fmt.Scan(&first_name)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&last_name)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter your password: ")
		fmt.Scan(&password)

		counter := exec(first_name, last_name, email, password)
		if counter != 0 {
			fmt.Println("you registered")
		} else {
			fmt.Println("such user already exists")
		}

	}

	fmt.Scan(&choise)
}

func exec(first_name string, last_name string, email string, password string) (counter int) {

	connStr := "user=postgres password=6858 dbname=productbd sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into users (first_name_user, last_name_user, email_user, password_user) values ($1, $2, $3, $4)", first_name, last_name, email, password)
	if err != nil {
		return
	}

	fmt.Println(result.RowsAffected())
	counter++
	return
}

func query(email string, password string) (counter int) {

	connStr := "user=postgres password=6858 dbname=productbd sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from users where email_user = $1 and password_user = $2", email, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		counter++
	}

	return
}

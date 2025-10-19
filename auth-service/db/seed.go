package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	var count int
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	username := os.Getenv("USERNAME")
	email := os.Getenv("EMAIL")
	role := "admin"
	pass := os.Getenv("PASS")

	err = DB.QueryRow("SELECT COUNT(*) FROM users WHERE role = 'admin'").Scan(&count)
	if err != nil {
		log.Fatal("Unable to query for admin user: ", err)
	}

	if count > 0 {
		fmt.Println("Admin user already exists")
		fmt.Println("Continuing the program ma boi hehehe")
		return
	}

	if count <= 0 {

		hashPass, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		_, err = DB.Exec("INSERT INTO users(username, email, password, role) VALUES($1, $2, $3, $4)", username, email, hashPass, role)
		if err != nil {
			log.Fatal("Unable to insert admin user: ", err)
		}
	}

	fmt.Println("Admin user seeded successfully")
}

package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"interview.heeko/login/model"
)

// Database ...
func Database() *sql.DB {
	databaseFilename := os.Getenv("DATABASE_FILENAME")
	filePath := fmt.Sprintf("database/%s", databaseFilename)
	// os.Remove(filePath)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	sqlLiteDatabase := getDatabase()
	createUserTable(sqlLiteDatabase)
	return sqlLiteDatabase
}

// CreateUser ...
func CreateUser(username string, password string) (*model.User, error) {
	db := getDatabase()
	insertUserSQL := `INSERT INTO users (username, password) VALUES (?, ?)`
	statement, _ := db.Prepare(insertUserSQL)
	_, err := statement.Exec(username, password)
	if err != nil {
		log.Println(fmt.Sprintf("Unable to create user due to: %s", err.Error()))
		return nil, err
	}
	user := model.User{
		Username: username,
		Password: password}
	log.Println(fmt.Sprintf("User added %s: %s", username, password))
	return &user, nil
}

// GetUser ...
func GetUser(username string, password string) (*model.User, error) {
	db := getDatabase()
	getUserSQL := `SELECT username, password FROM users WHERE username = ? AND password = ?`
	rows, err := db.Query(getUserSQL, username, password)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var username string
		var password string
		rows.Scan(&username, &password)
		user := model.User{
			Username: username,
			Password: password}
		return &user, nil
	}
	return nil, errors.New("User not found")
}

func getDatabase() *sql.DB {
	databaseFilename := os.Getenv("DATABASE_FILENAME")
	filePath := fmt.Sprintf("database/%s", databaseFilename)
	sqlLiteDatabase, err := sql.Open("sqlite3", filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	return sqlLiteDatabase
}

func createUserTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
		"username" string PRIMARY KEY,
		"password" string
	)`
	statement, err := db.Prepare(createUserTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("User table created")
}

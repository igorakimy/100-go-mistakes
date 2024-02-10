package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Принимается имя источника данных и возвращается *sql.DB и ошибка

func createClient(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// Возвращается ошибка
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// Ответственность на обработку ошибок возлагается на вызывающую функцию
	db, err := createClient(os.Getenv("MYSQL_DATA_SOURCE_NAME"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)
}

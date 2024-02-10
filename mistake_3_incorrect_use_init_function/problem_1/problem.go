package main

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// Переменная среды
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	// Определяет связь DB с глобальной переменной db
	db = d
}

func main() {
	// Невозможность обработки ошибок по своему, вынужденная паника
	// Инициализация будет выполняться перед запуском тестов, что не всегда нужно
	// Присвоение глобальной переменной, вместо инкапсуляции, что имеет ряд недостатков
}

package database

import (
	"fmt"
	"server/connection"
)

func CreateDB(databaseName string) {
	db := connection.Connect()
	sql := fmt.Sprintf("CREATE DATABASE %s\n", databaseName)
	fmt.Printf("Creating %s...\n", databaseName)

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Printf("Failed to create %s: %v\n\n", databaseName, err)
	} else {
		fmt.Printf("%s is created successfully!\n\n", databaseName)
	}

	defer db.Close()
}

func CreateTable(databaseName string, tableName string, id string) {
	db := connection.Connect()
	createTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (%s BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL);", databaseName, tableName, id)
	fmt.Printf("Creating %s table...\n", tableName)
	table, err := db.Exec(createTable)
	if err != nil {
		fmt.Printf("Failed to create %s table: %v\n", tableName, err)
	} else {
		rowsAffected, _ := table.RowsAffected()
		fmt.Printf("%s table is created successfully! RowsAffected: %d\n", tableName, rowsAffected)
	}
	defer db.Close()
}

func AddColumn(databaseName string, tableName string, columnName string, content string) {
	db := connection.Connect()

	addCol := fmt.Sprintf("ALTER TABLE %s.%s ADD %s %s;", databaseName, tableName, columnName, content)
	fmt.Printf("Creating %s %s column...\n", tableName, columnName)
	col, err := db.Exec(addCol)
	if err != nil {
		fmt.Printf("Failed to add %s %s column: %v\n", tableName, columnName, err)
	} else {
		rowsAffected, _ := col.RowsAffected()
		fmt.Printf("%s %s column is created successfully! RowsAffected: %d\n", tableName, columnName, rowsAffected)
	}

	defer db.Close()
}

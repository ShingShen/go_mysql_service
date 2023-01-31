package userEntity

import (
	"fmt"
	"server/connection"
)

func CreateUserTable() {
	db := connection.Connect()
	sql := `CREATE TABLE IF NOT EXISTS user (
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);`

	// sql := `DELIMITER $$
	// DROP PROCEDURE IF EXISTS addColumnToTable $$
	// CREATE PROCEDURE addColumnToTable()
	// BEGIN
	// IF NOT EXISTS( (SELECT * FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=DATABASE() AND COLUMN_NAME='usernamea' AND TABLE_NAME='user') )
	// THEN ALTER TABLE go_db.user ADD sale_person_designationzzz varchar(255);
	// END IF;
	// END $$
	// CALL addColumnToTable() $$
	// DELIMITER ;`

	fmt.Println("create user table...")

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Printf("create user table failed: %v", err)
		return
	}
	fmt.Println("create user table successfully!")

	defer db.Close()
}

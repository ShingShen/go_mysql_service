package userService

import (
	"fmt"
	"server/connection"
)

type User struct {
	ID          int    `json:"id"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

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

func CreateUser(user string, password string) {
	db := connection.Connect()
	res, err := db.Exec("INSERT INTO user(username, password) values(?,?)", user, password)
	if err != nil {
		fmt.Printf("Insert user data failed, err: %v", err)
		return
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Get insert id failed, err: %v", err)
		return
	}
	fmt.Println("Insert data id", lastInsertId)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v", err)
		return
	}
	fmt.Println("Affected rows: ", rowsAffected)
	defer db.Close()
}

func UpdateUser(id int, user string, password string) {
	db := connection.Connect()
	res, err := db.Exec("UPDATE user SET username=?, password=? WHERE id=?", user, password, id)
	if err != nil {
		fmt.Printf("Update user data failed, err: %v", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v", err)
		return
	}
	fmt.Println("Affected rows: ", rowsAffected)
	defer db.Close()
}

func GetUser(id int) (*User, error) {
	db := connection.Connect()
	defer db.Close()
	var user User
	res := db.QueryRow("SELECT * FROM user WHERE id=?", id)

	err := res.Scan(&user.ID, &user.UserName, &user.Password, &user.CreatedTime, &user.UpdatedTime)
	if err != nil {
		return &User{}, err
	} else {
		return &user, nil
	}
}

func GetAllUsers() []map[string]interface{} {
	db := connection.Connect()
	defer db.Close()
	res, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Printf("Failed to query users, err: %v", err)
	}

	columns, err := res.Columns()
	if err != nil {
		fmt.Printf("Failed to query users' column, err: %v", err)
	}
	fmt.Println("users' columns: ", columns)

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for res.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		res.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData
}

func DeleteUser(id int) {
	db := connection.Connect()
	res, err := db.Exec("DELETE FROM user WHERE id=?", id)
	if err != nil {
		fmt.Printf("Delete user data failed, err: %v", err)
		return
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Get delete id failed, err: %v", err)
		return
	}
	fmt.Println("Delete data id", lastInsertId)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err: %v", err)
		return
	}
	fmt.Println("Affected rows: ", rowsAffected)
	defer db.Close()
}

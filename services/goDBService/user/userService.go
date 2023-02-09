package userService

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"server/connection"
)

type User struct {
	UserId       uint64  `json:"user_id"`
	UserAccount  *string `json:"user_account"`
	UserPassword string  `json:"user_password"`
	Email        *string `json:"email"`
	Phone        *string `json:"phone"`
	CreatedTime  string  `json:"created_time"`
	UpdatedTime  *string `json:"updated_time"`
}

func CreateUser(userAccount string, userPassword string) string {
	var user User

	db := connection.Connect()
	defer db.Close()

	hash := sha256.New()
	hash.Write([]byte(userPassword))
	encodeUserPassword := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	// Checking duplicated user account
	checkUserAccountSql := `SELECT user_account FROM go_db.user WHERE user_account=?;`
	checkUserAccountRes := db.QueryRow(checkUserAccountSql, userAccount)
	checkUserAccountResErr := checkUserAccountRes.Scan(&user.UserAccount)
	if checkUserAccountResErr != nil {
		fmt.Println("Failed to check user account, Error: ", checkUserAccountResErr)
	}
	if user.UserAccount != nil {
		fmt.Println("The user account exists: ", *user.UserAccount)
		return "0"
	} else {
		fmt.Println("Do not have the user account, create new user account...")
	}

	// Creating new user account
	createUserSql := `INSERT INTO go_db.user(user_account, user_password) values(?,?)`
	createUserSqlRes, err := db.Exec(createUserSql, userAccount, encodeUserPassword)
	if err != nil {
		fmt.Printf("Failed to insert user data, err: %v", err)
	}
	lastInsertId, err := createUserSqlRes.LastInsertId()
	if err != nil {
		fmt.Printf("Failed to get inserted user id, err: %v", err)
	}
	fmt.Println("Inserted user id", lastInsertId)
	rowsAffected, err := createUserSqlRes.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get createUser affected rows, err: %v", err)
	}
	fmt.Println("CreateUser Affected rows: ", rowsAffected)
	return "-1"
}

func UpdateUser(userId uint64, email string, phone string) {
	db := connection.Connect()

	sql := `UPDATE go_db.user SET 
			email=?, 
			phone=? 
			WHERE user_id=?;`
	res, err := db.Exec(sql, email, phone, userId)
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

func GetUser(userId uint64) (*User, error) {
	db := connection.Connect()
	defer db.Close()
	var user User
	res := db.QueryRow("SELECT * FROM go_db.user WHERE user_id=?", userId)

	err := res.Scan(
		&user.UserId,
		&user.UserAccount,
		&user.UserPassword,
		&user.Email,
		&user.Phone,
		&user.CreatedTime,
		&user.UpdatedTime,
	)
	if err != nil {
		return &User{}, err
	} else {
		return &user, nil
	}
}

func GetAllUsers() []map[string]interface{} {
	db := connection.Connect()
	defer db.Close()

	sql := `SELECT * FROM go_db.user;`
	res, err := db.Query(sql)
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

func DeleteUser(userId uint64) {
	db := connection.Connect()

	sql := `DELETE FROM go_db.user WHERE id=?;`
	res, err := db.Exec(sql, userId)
	if err != nil {
		fmt.Printf("Failed to delete user data, err: %v", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v", err)
		return
	}
	fmt.Println("Affected rows: ", rowsAffected)
	defer db.Close()
}

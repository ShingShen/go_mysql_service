package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {
	const (
		Driver   string = "mysql"
		Database string = "go_db"
		IP       string = "0.0.0.0"
		Port     int    = 3306
		UserName string = "root"
		Password string = "password"
	)

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, IP, Port, Database)

	db, err := sql.Open(Driver, conn)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("successfully connected to database.")
	return db
}
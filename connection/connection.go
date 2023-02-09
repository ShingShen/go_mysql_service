package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {
	const (
		Driver string = "mysql"
		// Database string = "go_db"
		IP   string = "127.0.0.1"
		Port int    = 3306
		// UserName string = "rdrd"
		UserName string = "root"
		// Password string = "@Rdrdrd123"
		Password string = "root"
	)

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", UserName, Password, IP, Port)

	db, err := sql.Open(Driver, conn)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("successfully connected to database.")
	return db
}

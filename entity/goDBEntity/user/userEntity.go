package userEntity

import "server/utilities/database"

func CreateUserTable() {
	database.CreateTable("go_db", "user", "user_id")
	database.AddColumn("go_db", "user", "user_account", "VARCHAR(64)")
	database.AddColumn("go_db", "user", "user_password", "VARCHAR(64)")
	database.AddColumn("go_db", "user", "email", "VARCHAR(64)")
	database.AddColumn("go_db", "user", "phone", "VARCHAR(64)")
	database.AddColumn("go_db", "user", "created_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP")
	database.AddColumn("go_db", "user", "updated_time", "TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP")
}

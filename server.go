package main

import (
	"server/router"
	"server/utilities/database"

	goDBUserEntity "server/entity/goDBEntity/user"
)

func main() {
	database.CreateDB("go_db")
	goDBUserEntity.CreateUserTable()
	router.Handler()
}

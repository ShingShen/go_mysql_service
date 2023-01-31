package main

import (
	userEntity "server/entity/user"
	router "server/router"
)

func main() {
	userEntity.CreateUserTable()
	router.Handler()
}

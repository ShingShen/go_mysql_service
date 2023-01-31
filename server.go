package main

import (
	"fmt"
	"net/http"
	"server/web"

	router "server/router"
	userService "server/services/user"
)

func main() {
	http.HandleFunc("/", web.HomePage)
	userService.CreateUserTable()
	router.User()
	fmt.Println("Listening on port 3301......")
	http.ListenAndServe(":3301", nil)
}

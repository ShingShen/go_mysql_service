package router

import (
	"net/http"

	userController "server/controllers/user"
)

func User() {
	http.HandleFunc("/user/create", userController.Create)
	http.HandleFunc("/user/update", userController.Update)
	http.HandleFunc("/user/get", userController.GetUser)
	http.HandleFunc("/user/get_all", userController.GetAllUsers)
	http.HandleFunc("/user/delete", userController.Delete)
}

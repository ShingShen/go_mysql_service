package router

import (
	"fmt"
	"net/http"
	"time"

	userController "server/controllers/user"
)

func Handler() {
	mux := http.NewServeMux()

	// user
	mux.HandleFunc("/user/create", userController.Create)
	mux.HandleFunc("/user/update", userController.Update)
	mux.HandleFunc("/user/get", userController.GetUser)
	mux.HandleFunc("/user/get_all", userController.GetAllUsers)
	mux.HandleFunc("/user/delete", userController.Delete)

	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Addr:         ":3301",
	}

	fmt.Println("Listening on port 3301...")
	server.ListenAndServe()
}

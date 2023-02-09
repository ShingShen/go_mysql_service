package router

import (
	"fmt"
	"net/http"
	"time"

	goDBUserController "server/controllers/goDBController/user"
)

func Handler() {
	mux := http.NewServeMux()

	// go_db
	// user
	mux.HandleFunc("/user/create", goDBUserController.Create)
	mux.HandleFunc("/user/update", goDBUserController.Update)
	mux.HandleFunc("/user/get", goDBUserController.GetUser)
	mux.HandleFunc("/user/get_all", goDBUserController.GetAllUsers)
	mux.HandleFunc("/user/delete", goDBUserController.Delete)

	server := &http.Server{
		Handler:      mux,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Addr:         ":3301",
	}

	fmt.Printf("Listening on port %s......\n", server.Addr)
	server.ListenAndServe()
}

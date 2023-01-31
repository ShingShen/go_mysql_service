package userController

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	userService "server/services/user"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)
		userService.CreateUser(
			res["username"].(string),
			res["password"].(string),
		)
		fmt.Fprintf(w, "User created successfully!")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method == "PUT" {
		var res map[string]interface{}
		body, _ := ioutil.ReadAll(r.Body)
		userId, _ := strconv.Atoi(r.URL.Query().Get("id"))
		json.Unmarshal(body, &res)
		userService.UpdateUser(
			userId,
			res["username"].(string),
			res["password"].(string),
		)
		fmt.Fprintf(w, "User updated successfully!")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method == "GET" {
		userId, _ := strconv.Atoi(r.URL.Query().Get("id"))
		userData, _ := userService.GetUser(userId)
		userJsonData, _ := json.Marshal(userData)
		w.Write([]byte(userJsonData))
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method == "GET" {
		usersData := userService.GetAllUsers()
		usersJsonData, _ := json.Marshal(usersData)
		w.Write([]byte(usersJsonData))
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		userId, _ := strconv.Atoi(r.URL.Query().Get("id"))
		userService.DeleteUser(userId)
		fmt.Fprintf(w, "User deleted!")
	}
}

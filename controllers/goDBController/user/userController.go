package userController

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	goDBUserService "server/services/goDBService/user"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		var res map[string]interface{}
		json.Unmarshal(body, &res)
		goDBUserService.CreateUser(
			res["username"].(string),
			res["password"].(string),
		)
		fmt.Fprintf(w, "User created successfully!")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	if r.Method == "PUT" {
		var res map[string]interface{}
		body, _ := ioutil.ReadAll(r.Body)
		userId, _ := strconv.Atoi(r.URL.Query().Get("id"))
		json.Unmarshal(body, &res)
		goDBUserService.UpdateUser(
			uint64(userId),
			res["username"].(string),
			res["password"].(string),
		)
		fmt.Fprintf(w, "User updated successfully!")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	if r.Method == "GET" {
		userId, _ := strconv.Atoi(r.URL.Query().Get("id"))
		userData, _ := goDBUserService.GetUser(uint64(userId))
		userJsonData, _ := json.Marshal(userData)
		w.Write([]byte(userJsonData))
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	if r.Method == "GET" {
		usersData := goDBUserService.GetAllUsers()
		usersJsonData, _ := json.Marshal(usersData)
		w.Write([]byte(usersJsonData))
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	if r.Method == "DELETE" {
		userId, _ := strconv.Atoi(r.URL.Query().Get("id"))
		goDBUserService.DeleteUser(uint64(userId))
		fmt.Fprintf(w, "User deleted!")
	}
}

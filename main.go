package main

import (
	"app"
	"conf"
	"mydb"
	"net/http"
)

func main() {
	mydb.InitDB(conf.App.DBAddress)

	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/register", app.Register)

	http.HandleFunc("/users", app.ListUsers)

	http.ListenAndServe(":8888", nil)
}

package main

import (
	"crud/controllers"
	"crud/db"
	"fmt"
	"log"
	"net/http"

	
)


func main() {
	//connnecting to the db
	fmt.Println("connecting to db...")
	db.ConnectDb()
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/get", controllers.Get)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	//starting the server
	fmt.Println("starting the server...")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}

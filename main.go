package main

import (
	"CRUD-WEB/config"
	"CRUD-WEB/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	//databse connection
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server Runing On Port 8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"CRUD-WEB/config"
	"CRUD-WEB/controllers/categorycontroller"
	"CRUD-WEB/controllers/homecontroller"
	"CRUD-WEB/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	//databse connection
	config.ConnectDB()

	//Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//category
	http.HandleFunc("/categorys", categorycontroller.Index)
	http.HandleFunc("/categorys/add", categorycontroller.Add)
	http.HandleFunc("/categorys/edit", categorycontroller.Edit)
	http.HandleFunc("/categorys/delete", categorycontroller.Delete)

	//product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server Runing On Port 8080")
	http.ListenAndServe(":8080", nil)
}

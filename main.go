package main

import (
	"log"
	"net/http"
	"web_native/config"
	"web_native/controllers/categoriescontroller"
	"web_native/controllers/homecontroller"
	"web_native/controllers/productscontroller"
)

func main() {

	config.ConnectDB()

	// homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)	



	// products
	http.HandleFunc("/products", productscontroller.Index)
	http.HandleFunc("/products/add", productscontroller.Add)
	http.HandleFunc("/products/detail", productscontroller.Detail)
	http.HandleFunc("/products/edit", productscontroller.Edit)
	http.HandleFunc("/products/delete", productscontroller.Delete)

	log.Println("server running port 8080")
	http.ListenAndServe(":8080", nil)
}
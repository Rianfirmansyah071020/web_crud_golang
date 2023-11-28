package productscontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){

	temp,err := template.ParseFiles("views/products/index.html")
	
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}


func Add(w http.ResponseWriter, r *http.Request){

	temp,err := template.ParseFiles("views/products/add.html")
	
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}


func Edit(w http.ResponseWriter, r *http.Request){

	temp,err := template.ParseFiles("views/products/edit.html")
	
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}


func Delete(w http.ResponseWriter, r *http.Request){

	temp,err := template.ParseFiles("views/products/delete.html")
	
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
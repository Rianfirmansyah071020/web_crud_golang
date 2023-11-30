package productscontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){

	products := productmodel.GetAll()

	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/products/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}


func Add(w http.ResponseWriter, r *http.Request){

	temp,err := template.ParseFiles("views/products/add.html")
	
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}


func Detail(w http.ResponseWriter, r *http.Request){

	temp,err := template.ParseFiles("views/products/detail.html")
	
	if err != nil {
		panic(err)
	}
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
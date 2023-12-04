package productscontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"web_native/entities"
	"web_native/models/categorymodel"
	"web_native/models/productmodel"
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

	if(r.Method == "GET"){ 

		temp, err := template.ParseFiles("views/products/create.html")
		
		if err != nil {
			panic(err)
		}
categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}


		temp.Execute(w, data)


	}


	if(r.Method == "POST"){ 

		var product entities.Product
		
		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		


		if oke := productmodel.Create(product); !oke {
			temp, _ := template.ParseFiles("views/products/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/products/add", http.StatusSeeOther)

	}


	
}


func Detail(w http.ResponseWriter, r *http.Request){

	// temp,err := template.ParseFiles("views/products/detail.html")
	
	// if err != nil {
	// 	panic(err)
	// }
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
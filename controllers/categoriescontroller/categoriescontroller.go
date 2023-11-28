package categoriescontroller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
	"web_native/entities"
	"web_native/models/categorymodel"
)

func Index(w http.ResponseWriter, r *http.Request) { 
	
	categories := categorymodel.GetAll()

	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/categories/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)



}


func Add(w http.ResponseWriter, r *http.Request) { 

	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/categories/create.html")

		if(err != nil){
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" { 

		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if oke := categorymodel.Create(category); !oke {
			temp, _ := template.ParseFiles("views/categories/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories/add", http.StatusSeeOther)
	}

	
}



func Edit(w http.ResponseWriter, r *http.Request) { 

	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")

		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		data := map[string]any{
			"category": categorymodel.GetById(id),
		}		

		temp.Execute(w, data)
	}
}


func Delete(w http.ResponseWriter, r *http.Request) { 

	temp, err := template.ParseFiles("views/categories/delete.html")

	if(err != nil){
		panic(err)
	}

	temp.Execute(w, nil)
}
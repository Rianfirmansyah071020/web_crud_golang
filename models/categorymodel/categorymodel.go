package categorymodel

import (
	"web_native/config"
	"web_native/entities"
)

func GetAll()[]entities.Category {

  rows, err :=	config.DB.Query(`SELECT * FROM categories`)

  if err != nil {
  	panic(err)
  }

  defer rows.Close()

  var categories []entities.Category
  for rows.Next() {
  	var category entities.Category
   err :=	rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

   if err != nil {
   	panic(err)
   }

  	categories = append(categories, category)

}
return categories

}



func Create(category entities.Category) bool { 
  result, err := config.DB.Exec(`INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?)`, category.Name, category.CreatedAt, category.UpdatedAt)

  if err != nil {
    panic(err)
  }

  lastInsertId, err := result.LastInsertId()

  if err != nil {
    panic(err)
  }

  return lastInsertId > 0
}


func GetById(id int) entities.Category {

  rows, err := config.DB.Query(`SELECT * FROM categories WHERE id = ?`, id)

    if err != nil {
      panic(err)
    }

    defer rows.Close()

    var category entities.Category
    for rows.Next() {
    err :=	rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

    if err != nil {
      panic(err)
    }

  }
    return category
}
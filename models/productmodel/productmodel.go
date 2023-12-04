package productmodel

import (
	"web_native/config"
	"web_native/entities"
)

func GetAll() []entities.Product {

	rows, err := config.DB.Query(`SELECT 
	products.id,
	products.name,
	Categories.name as category_name,
	products.stock,
	products.description,
	products.created_at,
	products.updated_at	
	FROM products
	JOIN categories ON products.category_id = categories.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		err := rows.Scan(&product.Id, 
			&product.Name, 
			&product.Category.Name,
			&product.Stock, 
			&product.Description,
			&product.CreatedAt, 
			&product.UpdatedAt)

		if err != nil {
			panic(err)
		}

		products = append(products, product)

	}
	return products

}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`INSERT INTO products (name, created_at, updated_at) VALUES (?, ?, ?)`, product.Name, product.CreatedAt, product.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func GetById(id int) entities.Product {

	rows, err := config.DB.Query(`SELECT * FROM products WHERE id = ?`, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var product entities.Product
	for rows.Next() {
		err := rows.Scan(product.Id, product.Name, product.CreatedAt, product.UpdatedAt)

		if err != nil {
			panic(err)
		}

	}
	return product
}

func Update(id int, product entities.Product) bool {

	result, err := config.DB.Exec(`UPDATE products SET name = ?, updated_at = ? WHERE id = ?`, product.Name, product.UpdatedAt, id)

	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}

func Delete(id int) bool {

	result, err := config.DB.Exec(`DELETE FROM products WHERE id = ?`, id)

	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rowsAffected > 0

}
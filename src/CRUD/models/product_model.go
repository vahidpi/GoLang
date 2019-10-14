package models

import (
	"CRUD/entities"
	"database/sql"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product")
	if err != nil {
		return nil, err
	}
	products := []entities.Product{}
	for rows.Next() {
		var id int64
		var name string
		var price float32
		var quantity int
		var status bool
		err2 := rows.Scan(&id, &name, &price, &quantity, &status)
		if err2 != nil {
			return nil, err2
		}
		product := entities.Product{id, name, price, quantity, status}
		products = append(products, product)
	}
	return products, nil
}

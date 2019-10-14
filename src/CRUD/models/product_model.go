package models

import (
	"CRUD/entities"
	"database/sql"
)

//ProductModel is bublic
type ProductModel struct {
	Db *sql.DB
}

//Create is public
func (productModel ProductModel) Create(product *entities.Product) error {
	result, err := productModel.Db.Exec("insert into product(name,price,quantity,status) values(?,?,?,?)", product.Name, product.Price, product.Quantity, product.Status)

	if err != nil {
		return err
	}
	product.ID, _ = result.LastInsertId()
	return nil
}

//Find is public
func (productModel ProductModel) Find(id int64) (entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where id = ?", id)
	if err != nil {
		return entities.Product{}, err
	}
	product := entities.Product{}
	for rows.Next() {
		var id int64
		var name string
		var price float32
		var quantity int
		var status bool
		err2 := rows.Scan(&id, &name, &price, &quantity, &status)
		if err2 != nil {
			return entities.Product{}, err2
		}
		product = entities.Product{id, name, price, quantity, status}
	}
	return product, nil
}

//FindAll is public
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

package main

import (
	"CRUD/config"
	"CRUD/entities"
	"CRUD/models"
	"fmt"
)

func main() {
	demoDelete()
	demoCallFindAll()

}

func demoDelete() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		rows, err := productModel.Delete(3)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("Done.")
			}
		}
	}
}

func demoUpdate() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			ID:       5,
			Name:     "def",
			Price:    44.5,
			Quantity: 4,
			Status:   false,
		}
		rows, err := productModel.Update(product)
		if err != nil {
			fmt.Println(err)
		} else {
			if rows > 0 {
				fmt.Println("Done.")
			}
		}
	}
}

func demoInsert() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			Name:     "abc",
			Price:    4.5,
			Quantity: 777,
			Status:   true,
		}
		err := productModel.Create(&product)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Lastest Id: ", product.ID)
		}
	}
}

func demoCallFind() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err := productModel.Find(1)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(product)
			fmt.Println("Product Info")
			fmt.Println("Id:", product.ID)
			fmt.Println("Name:", product.Name)
			fmt.Println("Price:", product.Price)
			fmt.Println("Quantity:", product.Quantity)
			fmt.Println("Status:", product.Status)

		}
	}
}

func demoCallFindAll() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err := productModel.FindAll()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(products)
			fmt.Println("Product List")
			for _, product := range products {
				fmt.Println("Id:", product.ID)
				fmt.Println("Name:", product.Name)
				fmt.Println("Price:", product.Price)
				fmt.Println("Quantity:", product.Quantity)
				fmt.Println("Status:", product.Status)
				fmt.Println("----------------------")
			}
		}
	}

}

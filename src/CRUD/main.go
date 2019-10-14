package main

import (
	"CRUD/config"
	"CRUD/models"
	"fmt"
)

func main() {
	demo1CallFindAll()
}

func demo1CallFindAll() {
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

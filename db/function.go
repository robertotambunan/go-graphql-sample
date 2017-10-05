package db

import (
	"fmt"
	"log"

	"github.com/tokopedia/project2/types"
)

func GetProductData() []types.Product {
	db := OpenConnection()

	defer CloseConnection(db)

	rows, err := db.Query("SELECT * FROM product")

	var products []types.Product

	if err != nil {
		return products
	}

	for rows.Next() {
		var product_id int
		var product_name string
		var product_price string
		var product_city string
		err = rows.Scan(&product_id, &product_name, &product_price, &product_city)
		if err != nil {
			log.Println("Cannot Query MySQL Database")
		}
		products = append(products, types.Product{
			ProductID:    product_id,
			ProductCity:  product_city,
			ProductName:  product_name,
			ProductPrice: product_price,
		})
	}

	return products

}

func AddProductData(product_name string, product_city string, product_price string) bool {
	db := OpenConnection()
	defer CloseConnection(db)

	stmt, err := db.Prepare("INSERT product set product_name=?, product_city=?, product_price=?")
	if err != nil {
		fmt.Println("Error in Query Add Product")
		return false
	}

	_, errs := stmt.Exec(product_name, product_city, product_price)
	if errs != nil {
		fmt.Println("Error in Exec Query Add Product")
		return false
	}

	return true

}

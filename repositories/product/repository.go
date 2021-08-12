package product

import (
	models "goShop/models"
	db "goShop/shared/db"
)

func ListProducts() []models.Product {
	products := []models.Product{}
	dbDataSet, err := db.RunQuery("SELECT name, description, price, amount FROM product")
	if err != nil {
		panic(err)
	}

	for dbDataSet.Next() {
		var name, description string
		var price float64
		var amount int32

		dbDataSet.Scan(&name, &description, &price, &amount)
		eachProduct := models.Product{
			Name:        name,
			Description: description,
			Price:       price,
			Amount:      amount,
		}
		products = append(products, eachProduct)
	}

	return products
}

func Save(p models.Product) {
	conn, query, err := db.Prepare("INSERT INTO product (name, description, price, amount) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = query.Exec(p.Name, p.Description, p.Price, p.Amount)
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()
}

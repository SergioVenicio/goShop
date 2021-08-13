package product

import (
	"fmt"
	models "goShop/models"
	db "goShop/shared/db"
)

func getSqlFilePath(fileName string) string {
	return "repositories/products/sql/" + fileName + ".sql"
}

func ListProducts() []models.Product {
	products := []models.Product{}
	dbDataSet, err := db.RunQuery(db.LoadQuery(getSqlFilePath("list_products")))
	if err != nil {
		panic(err)
	}

	for dbDataSet.Next() {
		var id int
		var name, description string
		var price float64
		var amount int

		dbDataSet.Scan(&id, &name, &description, &price, &amount)
		eachProduct := models.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Amount:      amount,
		}
		products = append(products, eachProduct)
	}

	return products
}

func GetById(id int) models.Product {
	var product = models.Product{}
	query := fmt.Sprintf(db.LoadQuery(getSqlFilePath("get_by_id")), id)
	dbDataSet, err := db.RunQuery(query)
	if err != nil {
		panic(err)
	}

	for dbDataSet.Next() {
		var id int
		var name, description string
		var price float64
		var amount int

		dbDataSet.Scan(&id, &name, &description, &price, &amount)
		product = models.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Amount:      amount,
		}
	}

	return product
}

func Edit(p models.Product) {
	conn, query, err := db.Prepare(db.LoadQuery(getSqlFilePath("update_product")))
	if err != nil {
		panic(err.Error())
	}

	_, err = query.Exec(p.Id, p.Name, p.Description, p.Price, p.Amount)
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()
}

func Save(p models.Product) {
	conn, query, err := db.Prepare(db.LoadQuery(getSqlFilePath("insert_product")))
	if err != nil {
		panic(err.Error())
	}

	_, err = query.Exec(p.Name, p.Description, p.Price, p.Amount)
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()
}

func DeleteById(id int) {
	conn, query, err := db.Prepare(db.LoadQuery(getSqlFilePath("delete_product")))
	if err != nil {
		panic(err.Error())
	}

	_, err = query.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()
}

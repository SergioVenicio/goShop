package routes

import (
	"net/http"

	"goShop/controllers/products"
)

func LoadRoutes() {
	http.HandleFunc("/", products.Index)
	http.HandleFunc("/product/new", products.New)
	http.ListenAndServe(":8000", nil)
}

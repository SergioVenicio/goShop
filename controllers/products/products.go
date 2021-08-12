package products

import (
	"html/template"
	"net/http"
	"strconv"

	"goShop/models"
	repository "goShop/repositories/product"
)

var templateFolder = template.Must(template.ParseGlob("templates/products/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := repository.ListProducts()
	templateFolder.ExecuteTemplate(w, "index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		amount, _ := strconv.Atoi(r.FormValue("amount"))

		if name == "" || description == "" || price == 0 || amount == 0 {
			templateFolder.ExecuteTemplate(w, "form_error", nil)
		}

		product := models.Product{
			Name:        name,
			Description: description,
			Price:       price,
			Amount:      int32(amount),
		}
		repository.Save(product)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	templateFolder.ExecuteTemplate(w, "new", nil)
}

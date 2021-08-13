package products

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"goShop/models"
	repository "goShop/repositories/products"
)

var templateFolder = template.Must(template.ParseGlob("templates/*/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := repository.ListProducts()
	templateFolder.ExecuteTemplate(w, "product_index", products)
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
			Amount:      amount,
		}
		repository.Save(product)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	templateFolder.ExecuteTemplate(w, "product_new", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			panic(err.Error())
		}

		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		amount, _ := strconv.Atoi(r.FormValue("amount"))

		if name == "" || description == "" || price == 0 || amount == 0 {
			templateFolder.ExecuteTemplate(w, "form_error", nil)
		}

		product := models.Product{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Amount:      amount,
		}
		repository.Edit(product)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	urlId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(urlId)
	if urlId == "" || err != nil {
		http.Error(w, "Invalid id!", http.StatusBadRequest)
	}

	product := repository.GetById(id)
	templateFolder.ExecuteTemplate(w, "product_edit", product)
}

type DeleteResponse struct {
	Message string
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(DeleteResponse{
			Message: "Http method not allowed!",
		})
		return
	}

	urlId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(urlId)
	if urlId == "" || err != nil {
		http.Error(w, "Invalid id!", http.StatusBadRequest)
	}

	repository.DeleteById(id)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(DeleteResponse{
		Message: "Ok",
	})
}

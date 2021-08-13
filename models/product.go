package models

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func New(name string, description string, price float64, amount int) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Amount:      amount,
	}
}

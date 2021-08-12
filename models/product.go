package models

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int32
}

func New(name string, description string, price float64, amount int32) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Amount:      amount,
	}
}

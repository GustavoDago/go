package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products = []Product{
	{ID: 1, Name: "Product1", Price: 10.99, Description: "Description1", Category: "Category1"},
	{ID: 2, Name: "Product2", Price: 20.99, Description: "Description2", Category: "Category2"},
}

func (p Product) Save() {
	Products = append(Products, p)
}

func GetAll() {
	for _, p := range Products {
		fmt.Println("ID:", p.ID, "Name:", p.Name, "Price:", p.Price, "Description:", p.Description, "Category:", p.Category)	
	}
}

func getById(id int) Product {
	respuesta := Product{}
	for _, p := range Products {
		if p.ID == id{
			respuesta = p
		}
	}
	return respuesta
}

func main() {
	prod := Product{
		ID: 3, 
		Name: "Product3", 
		Price: 30.99, 
		Description: "Description3", 
		Category: "Category3",
	}

	prod.Save()

	GetAll()

	fmt.Println(getById(2))


}

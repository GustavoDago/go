package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	CodeValue string `json:"codevalue"`
	Is_Published bool `json:"is_published"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}

var listaDeProductos = []Product{}

func main() {
	router := gin.Default()
	loadProductos("./products.json",&listaDeProductos)

	router.GET("/products/search", SearchProducts())
	router.Run()
}

func loadProductos(path string, list *[]Product)  {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}

func SearchProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		consulta := ctx.Query("priceGT")
		priceGt, err := strconv.ParseFloat(consulta,64)
		if err != nil {
			ctx.JSON(400,gin.H{"error":"Formato inválido"})
		}


		list := []Product{}

		for _, product := range listaDeProductos {
			if product.Price > priceGt {
				list = append(list, product)
			}
		}
		if len(list)>0 {
			ctx.JSON(200,list)
			
		}else {
			ctx.JSON(404, gin.H{"msg":"No hay productos con precios mayores"})
		}

	}
}
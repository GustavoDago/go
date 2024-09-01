package main

import (
	"encoding/json"
	// "fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	CodeValue    string  `json:"codevalue"`
	Is_Published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

var listaDeProductos = []Product{}
var idVerificar int

func main() {
	router := gin.Default()
	loadProductos("./products.json", &listaDeProductos)

	router.GET("/products/search", SearchProducts())
	router.GET("/productsparams", func(ctx *gin.Context) {
		productoGenerado, err := leerParametros(ctx)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Formato inválido"})
			return
		}
		listaDeProductos = append(listaDeProductos, productoGenerado)
		ctx.JSON(200, productoGenerado)
	})
	router.GET("/products/:id", func(ctx *gin.Context) {
		idVerificar, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Formato de ID inválido"})
			return
		}
		productoTomado := tomarProducto(idVerificar)
		ctx.JSON(200, productoTomado)
	})
	router.Run()
}

func tomarProducto(idVerificar int) Product {
	productoTomado := Product{}
	for _, product := range listaDeProductos {
		if product.Id == idVerificar {
			productoTomado = product
			break
		}
	}
	return productoTomado
}

func loadProductos(path string, list *[]Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}

func leerParametros(ctx *gin.Context) (Product, error) {
	idConsultado := ctx.Query("id")
	idParseado, errorParseado := strconv.ParseInt(idConsultado, 10, 32)
	idVerificar = int(idParseado)
	if errorParseado != nil {
		ctx.JSON(400, gin.H{"error": "Formato inválido"})
		return Product{}, errorParseado
	}
	nameQueried := ctx.Query("name")

	quantityQueried := ctx.Query("quantity")
	quantityParsed, errorParsingQuantity := strconv.ParseInt(quantityQueried, 10, 32)

	if errorParsingQuantity != nil {
		ctx.JSON(400, gin.H{"error": "Formato Json inválido"})
		return Product{}, gin.Error{}
	}

	codeValueQueried := ctx.Query("code_value")

	isPublishedQueried := ctx.Query("is_published")

	isPublishedParsed, errorParsingIsPublished := strconv.ParseBool(isPublishedQueried)

	if errorParsingIsPublished != nil {
		ctx.JSON(400, gin.H{"error": "Formato Json inválido"})

	}

	expirationQueried := ctx.Query("expiration")

	priceQueried := ctx.Query("price")

	priceParsed, errorParsingQuery := strconv.ParseFloat(priceQueried, 64)

	if errorParsingQuery != nil {
		ctx.JSON(400, gin.H{"error": "Formato Json inválido"})

	}

	productoGenerado := Product{
		Id:           int(idParseado),
		Name:         nameQueried,
		Quantity:     int(quantityParsed),
		CodeValue:    codeValueQueried,
		Is_Published: isPublishedParsed,
		Expiration:   expirationQueried,
		Price:        priceParsed,
	}
	return productoGenerado, nil

}

func SearchProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		consulta := ctx.Query("priceGT")
		priceGt, err := strconv.ParseFloat(consulta, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Formato inválido"})
		}

		list := []Product{}

		for _, product := range listaDeProductos {
			if product.Price > priceGt {
				list = append(list, product)
			}
		}
		if len(list) > 0 {
			ctx.JSON(200, list)

		} else {
			ctx.JSON(404, gin.H{"msg": "No hay productos con precios mayores"})
		}

	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type producto struct {
	Id              string
	Nombre          string
	Precio          string
	Stock           float32
	Codigo          string
	Publicado       string
	FechaDeCreacion string
}

func main() {
	archivoJson, err := os.ReadFile("./productos.json")
	
	if err != nil {
		panic("Error en la lectura de nuestro archivo.")
	}
	var productos []producto
	
	if err := json.Unmarshal(archivoJson, &productos); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Lista de productos:")
	for _, p := range productos {
		fmt.Println(p)
	}

	router := gin.Default()

	router.GET("/productos",func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"producto":productos,
		})
	})

	router.Run(":8080")
}

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type persona struct {
	Nombre    string
	Apellido  string
	Edad      int
	Direccion string
	Telefono  string
	Activo    bool
}

func main() {
	router := gin.Default()
	jsonPer := `{"Nombre":"Juan", "Apellido":"Perez",
		"Edad":23,"Direccion":"dfg","Telefono": "1321113","Activo":true}`
	var p persona

	if err := json.Unmarshal([]byte(jsonPer), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)

	router.GET("/persona2", func(ctx *gin.Context) {
		panic("test")
		ctx.JSON(200, gin.H{
			"persona": p,
		})
	})
	router.GET("/persona", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"persona": p,
		})
	})
	router.Run(":8080")
}

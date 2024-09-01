# Consignas-Go-Web
En la rama main encontramos el proyecto completo de Go Web, en cada rama encontramos las respectivas consignas.

## Ejercicio: Swagger.io

Documentamos nuestra API utilizando **Swagger**, pero haremos su implementación en Go. Recuerden lo visto en clase, también les dejamos algunos tips extra:

- Para utilizar swaggo debemos instalarlo y, cada vez que busquemos actualizar la documentación, debemos correr el comando **swag init**
- Debemos añadir a nuestro entorno la variable HOST con la dirección de nuestra API.
- La ruta de la path a la página de swagger es por defecto: */swagger/index.html*.
- Un ejemplo de la documentación del método Post puede ser:

```go
// Post godoc
// @Summary      Create a new product
// @Description  Create a new product in repository
// @Tags         products
// @Produce      json
// @Param        token header string true "token"
// @Param        body body domain.Product true "Product"
// @Success      201 {object}  web.response
// @Failure      400 {object}  web.errorResponse
// @Failure      404 {object}  web.errorResponse
// @Router       /products [post]
func (h *productHandler) Post() gin.HandlerFunc {}
```

Para implementar Swagger utilizaremos **swaggo**.

```terminal
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/files
go get -u github.com/swaggo/swag/gin-swagger

export PATH=$PATH:$HOME/go/bin
```

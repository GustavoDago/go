# Especialización en Back End III

Mesa de trabajo

## Ejercitación grupal

Nivel de complejidad: medio 🔥🔥

### Problema

Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Por este motivo, necesitan un servidor que ejecute una API que les permita manipular los productos cargados de distintos clientes.
Siguiendo con el ejercicio en clase, vamos a implementar el paquete store para tener un mejor control sobre nuestros productos. Además, vamos a mejorar nuestras responses.

#### Ejercicio 1: Paquete store

En lugar de trabajar sobre una lista cargada en memoria, ahora tenemos que definir un paquete store que nos servirá como interfaz para modificar el archivo .json de productos. Este paquete debe estar dentro de una carpeta pkg (la misma contiene paquetes de uso general y utilidades para el consumo de la aplicación).
Debemos implementar funciones de inicialización (verificar que es posible leer el archivo y modificarlo), búsqueda (buscar un producto concreto por ID), modificación (actualizar campos de un producto por ID) y borrado (eliminar un producto por ID).
Este paquete se debe iniciar en el main.go de nuestra API y ser utilizado por Repository, implementando los respectivos métodos en las interfaces que definimos en el ejercicio anterior.

#### Ejercicio 2: Manejo de responses

Definimos un paquete web cuya función será definir un estándar para los responses de nuestra API. Dentro de él, tendremos funciones para las condiciones de éxito y otra para las de fallo. A estas funciones es buena idea pasarles por argumento un c *gin.Context. Con esto podemos acceder a los métodos como c.Json() y además trasladar el contexto a la ejecución de las funciones. Las estructuras son:

```go
type errorResponse struct {
   Status int `json:"status"`
   Code string `json:"code"`
   Message string `json:"message"`
}

type response struct {
   Data interface{} `json:"data"`
}
```

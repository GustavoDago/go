# Especialización en Back End III
## Mesa de trabajo  
### Ejercitación grupal
Nivel de complejidad: medio 🔥🔥
## Problema

Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Para poder hacer esto, necesitan un servidor que ejecute una API que les permita manipular los productos cargados desde distintos clientes. Los campos que conforman un producto son:

|Nombre|Tipo de dato|Descripción | Ejemplo|
|-|-|-|-|
|id|number|Identificador en conjunto de datos | 15|
|name|string|Nombre caracterizado | Cheese - St. Andre|
|quantity|number|Cantidad almacenada | 60|
|code_value|string|Código alfanumérico característico | S73191A|
|is_published|boolean|El producto se encuentra publicado o no |  True|
|expiration|string|Fecha de vencimiento | 12/04/2022|
|price|number|Precio del producto | 50.15|





#### IMPORTANTE

En el siguiente enlace, verás la solución del ejercicio anterior para que puedas continuar si no lograste terminar.

### Ejemplo de payload


```
{
   "name":"Chicken",
   "quantity": 1,
   "code_value":"AF429XD",
   "is_published": true,
   "expiration":"10/11/2009",
   "price": 10.50
}
```

### Ejercicio: Método PATCH
Añadir el método PATCH a nuestra API. Recordemos que este aplica modificaciones parciales a un recurso. Cuando realizamos un PATCH, solo debemos indicar los campos que vamos a modificar.
### Ejercicio: Método DELETE
Añadir el método DELETE a nuestra API. Recordemos que este elimina un recurso. Una vez eliminado, no debe ser posible acceder a él.
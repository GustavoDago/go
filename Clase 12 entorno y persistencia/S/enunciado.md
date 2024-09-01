# Especialización en Back End III

## Procesando peticiones  

### Ejercitación grupal

Nivel de complejidad: medio 🔥🔥

#### Problema

Un supermercado necesita un sistema para gestionar los productos frescos que tienen publicados en su página web. Para poder hacer esto, necesitan un servidor que ejecute una API que les permita manipular los productos cargados desde distintos clientes...

#### IMPORTANTE

En el siguiente enlace, verás la solución del ejercicio anterior para que puedas continuar si no lograste terminar.

#### Ejemplo de payload

```json
{
  "name":"Chicken",
   "quantity": 1,
   "code_value":"AF429XD",
   "is_published": true,
   "expiration":"10/11/2009",
   "price": 10.50
}
```

#### Ejercicio: Método PUT

Añadir el método PUT a nuestra API. Recordemos que este crea o reemplaza un recurso en su totalidad con el contenido en el request. Debemos validar los campos que se envían, tal como hicimos con el método POST. Seguimos aplicando los cambios sobre la lista cargada en memoria.

Para implementarlo, debemos:

1. Definir la ruta PUT /products

2. Validar que los campos obligatorios estén presentes

3. Actualizar el producto con el id correspondiente o crearlo si no existe

4. Retornar status 200 y el producto actualizado

5. Manejar posibles errores

Esto nos permitirá modificar un producto existente o crear uno nuevo a través de PUT.
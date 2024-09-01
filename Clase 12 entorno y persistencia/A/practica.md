# Especialización en Back End III  
## Práctica
### Ejercitación individual
Nivel de complejidad: medio 🔥🔥
#### Objetivo
El objetivo de esta guía práctica es que podamos afianzar los conceptos sobre variables de entorno y manipulación de archivos, vistos en esta clase.
#### Problema
Para esta ocasión, vamos a trabajar con lo que acabamos de aprender sobre variables de entorno y su uso.
### Ejercicio 1: Configurar main 
Vamos a necesitar el uso de variables entorno en nuestra nueva API, para eso vamos a crear dos variables de entorno en un fichero .env que contengan los nombres USER y PASSWORD.
Vamos a importar el paquete godotenv a través de:
`go get -u github.com/joho/godotenv`
Estas variables se deberán iniciar cada una con su respectiva palabra. Por ejemplo:
`.env`
`USER=NombreUsuario`
`PASSWORD=Contraseña`

### Ejercicio 2: Validar variables de entorno
Cada vez que iniciemos nuestro main, debemos validar que se cargue nuestro fichero .env, y lanzar un error si es que no se encuentra. Además, se deben tomar estas variables de entorno en nuestro main e imprimirlas por consola si es que las encuentra.

### Ejercicio 3: autoload  
Investigar y hacer uso del autoload que nos proporciona godotenv en nuestra misma API y anotar las ventajas y desventajas de hacer uso de la misma.
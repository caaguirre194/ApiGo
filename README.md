# ApiGo
Este proyecto es una **API REST** en donde se trabajó por parte del servidor con Golang mientras que del lado del cliente se utilizó Angular.js
## GO-ANGULAR:
 ![GitHub](/img/go-angular.png)
 
Using:
* [x] [Go](https://golang.org/) 
* [x] [Angular](https://angular.io/)
* [x] [Bootstrap](https://v4-alpha.getbootstrap.com/)

![Logo](/img/angular.png)
![Logo](/img/bootstrap.png)

# Despliegue de aplicación:
> ## Requisitos previos:
> - Tener instalado el sistema de gestión de bases de datos relacional MySQL.
	
![Logo](/img/mysql.png)

> - Crear una base de datos con las sigientes caracteristicas o bien modificar la variable *const* a conveniencia en main.go:

```go
		
	// Host.
	DbHost = "tcp(127.0.0.1:3306)"
	// Nombre de la Base de Datos.
	DbName = "sakila2"
	// Nombre de Usuario.
	DbUser = "root"
	// Contraseña.
	DbPass = "suerte123"
		
```	
> - Tener instalado un entorno de desarrollo web como lo es WampServer para Windows, el cual permite crear aplicaciones web con Apache2 y una base de datos MySQL.
 
![Logo](/img/wamp.png)

## Ejecución:
> - Ejecutar el documento main.go mediante el comando *go run main.go* en el respectivo directorio del servidor de la aplicación. 
> - Ejecutar el documento index.html mediante el entorno de desarrollo web previamente instalado con Apache2. Esto en su respectivo directorio como cliente de la aplicación.

## Visualización:
# Página de inicio:
![Logo](/img/page1.png)
# Agregar Usuario:
![Logo](/img/page2.png)
# Consultar usuario:
![Logo](/img/page3.png)

Autor:
*  [GitHub](https://github.com/caaguirre194)
	 @caaguirre194


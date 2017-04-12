package main

import (
	// Permite manejo de Logs.
	"log"
	// Permite conversiones de tipo string a otros tipos de datos basicos.
	"strconv"
	// Permite la conección a bases de datos.
	"database/sql"
	// Permit conectarse a bases de datos Mysql.
	_ "github.com/go-sql-driver/mysql"
	// Permite generar automaticamente las tablas de la base de datos SQL.
	"github.com/coopernurse/gorp"
	// Es un framekork web con alto rendimiento.
	"github.com/gin-gonic/gin"
)

// Constante de conexión con la Base de Datos.
const (
	// Host.
	DbHost = "tcp(127.0.0.1:3306)"
	// Nombre de la Base de Datos.
	DbName = "sakila"
	// Nombre de Usuario.
	DbUser = "root"
	// Contraseña.
	DbPass = "suerte123"
)

// User representa la Tabla de la Base de Datos Usuario.
type User struct {
	// ID - Primary Key de la Tabla.
	ID int64 `db:"id" json:"id"`
	// Nombre del Usuario.
	Firstname string `db:"firstname" json:"firstname"`
	// Aperllido del Usuario.
	Lastname string `db:"lastname" json:"lastname"`
}

// Cors establece la relacion con los controladores de la parte del cliente.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())
	// v1 Genera las rutas de peticiones a los controladores.
	v1 := r.Group("api/v1")
	{
		v1.GET("/users", getUsers)
		v1.GET("/users/:id", getUser)
		v1.POST("/users", insertUser)
		v1.PUT("/users/:id", updateUser)
		v1.DELETE("/users/:id", deleteUser)
		v1.OPTIONS("/users", optionsUser)     // POST
		v1.OPTIONS("/users/:id", optionsUser) // PUT, DELETE
	}
	r.Run(":8080")
}

// Variable que representa la Base de Datos.
var Dbmap = initDb()

// initDb es la función que permite conectarse a la Base de Datos.
func initDb() *gorp.DbMap {
	dsn := DbUser + ":" + DbPass + "@" + DbHost + "/" + DbName + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	checkErr(err, "sql.Open failed")
	Dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbmap.AddTableWithName(User{}, "User").SetKeys(true, "ID")
	err = Dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return Dbmap
}

// checkErr permite verificar errores en las funciones que retornan dos valores.
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// getUsers permite traer los usuarios de la Tabla User.
func getUsers(c *gin.Context) {
	var users []User
	_, err := Dbmap.Select(&users, "SELECT * FROM User")
	if err == nil {
		c.JSON(200, users)
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}
	// Verificar usuarios en ----> http://localhost:8080/api/v1/users
}

// getUser permite traer a un usuario de la tabla User.
func getUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	err := Dbmap.SelectOne(&user, "SELECT * FROM User WHERE id=?", id)
	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)
		content := &User{
			ID:        user_id,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
	// Verificar usuario en ----> http://localhost:8080/api/v1/users/1
}

// insertUser permite agregar un nuevo usuario en la Tabla User.
func insertUser(c *gin.Context) {
	var user User
	c.Bind(&user)
	if user.Firstname != "" && user.Lastname != "" {
		if insert, _ := Dbmap.Exec(`INSERT INTO User (firstname, lastname) VALUES (?, ?)`, user.Firstname, user.Lastname); insert != nil {
			user_id, err := insert.LastInsertId()
			if err == nil {
				content := &User{
					ID:        user_id,
					Firstname: user.Firstname,
					Lastname:  user.Lastname,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
}

// updateUser permite actualizar los datos de un usuario en la tabla User.
func updateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	err := Dbmap.SelectOne(&user, "SELECT * FROM User WHERE id=?", id)
	if err == nil {
		var json User
		c.Bind(&json)
		user_id, _ := strconv.ParseInt(id, 0, 64)
		user := User{
			ID:        user_id,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
		}
		if user.Firstname != "" && user.Lastname != "" {
			_, err = Dbmap.Update(&user)
			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}
		} else {
			c.JSON(422, gin.H{"error": "fields are empty"})
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

// deleteUser permite eliminar un campo de la tabla User.
func deleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	err := Dbmap.SelectOne(&user, "SELECT id FROM User WHERE id=?", id)
	if err == nil {
		_, err = Dbmap.Delete(&user)
		if err == nil {
			c.JSON(200, gin.H{"id #" + id: " Eliminado."})
		} else {
			checkErr(err, "Delete failed")
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

// optionsUser permite garantizar que las funcoines DELETE, POST y PUT tengan los permisos necesarios.
func optionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

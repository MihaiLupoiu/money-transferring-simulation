package main

import (
	"fmt"
	"math"
	"os"

	. "github.com/MihaiLupoiu/money-transferring-simulation/backend/libs/constants"
	. "github.com/MihaiLupoiu/money-transferring-simulation/backend/libs/util"
	"github.com/MihaiLupoiu/money-transferring-simulation/backend/models"
	"github.com/MihaiLupoiu/money-transferring-simulation/backend/queries"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/users", Post)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.PUT("/users/:id", Update)
		v1.DELETE("/users/:id", Delete)

		v1.GET("/kill", Kill)
		v1.GET("/pi", PiNumber)

		// TODO: split functions in seperate module.
		v1.GET("/balance/:id", user.GetBalance)
		v1.POST("/deposit/:id", user.AddDeposit)
		v1.POST("/transfer/:id", user.Transfer)
	}

	r.Run(":8080")
}

// Not necesary at the moment.
// func OptionsUser(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	c.Next()
// }

// TODO: Split DB queries and JSON return values.

func Post(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user models.User
	c.Bind(&user)

	if user.Firstname != "" && user.Lastname != "" && user.Mail != "" {
		db.Create(&user)
		c.JSON(201, gin.H{"success": user})
	} else {
		c.JSON(422, gin.H{"error": ErrorFieldsEmpty})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Donals\", \"mail\": \"jd@fake.com\"}" http://localhost:8080/api/v1/users
}

func GetUsers(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var users []models.User
	db.Find(&users)
	c.JSON(200, users)

	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Id != 0 {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": ErrorUserNotFound})
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

func Update(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Firstname != "" && user.Lastname != "" && user.Mail != "" {
		if user.Id != 0 {
			var newUser models.User
			c.Bind(&newUser)

			result := models.User{
				Id:        user.Id,
				Firstname: newUser.Firstname,
				Lastname:  newUser.Lastname,
				Mail:      newUser.Mail,
			}

			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, gin.H{"error": ErrorUserNotFound})
		}
	} else {
		c.JSON(422, gin.H{"error": ErrorFieldsEmpty})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\", \"mail\": \"tm@fake.com\" }" http://localhost:8080/api/v1/users/1
}

func Delete(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Id != 0 {
		db.Delete(&user)
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": ErrorUserNotFound})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func Kill(c *gin.Context) {
	fmt.Println("EXITING BRUTE FORCE!")
	os.Exit(-1)

	// curl -i -X GET http://localhost:8080/api/v1/kill
}

func PiNumber(c *gin.Context) {
	c.JSON(200, gin.H{"success": "Pi is" + fmt.Sprintf("%.10f", pi(10000))})
	// curl -i -X GET http://localhost:8080/api/v1/pi
}

// pi launches n goroutines to compute an
// approximation of pi.
func pi(n int) float64 {
	ch := make(chan float64)
	//defer close(ch)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

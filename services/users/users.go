package main

import (
	"github.com/MihaiLupoiu/money-transferring-simulation/queries"
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
		v1.POST("/users", user.Post)
		v1.GET("/users", user.GetUsers)
		v1.GET("/users/:id", user.GetUser)
		v1.PUT("/users/:id", user.Update)
		v1.DELETE("/users/:id", user.Delete)

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

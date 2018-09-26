package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/MihaiLupoiu/money-transferring-simulation/backend/libs/db"
	"github.com/MihaiLupoiu/money-transferring-simulation/backend/models/task"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", GetTasks)
			tasks.POST("/", CreateTask)
			tasks.PUT("/:id", UpdateTask)
			tasks.DELETE("/:id", DeleteTask)
		}

		v1.GET("/kill", Kill)
		v1.GET("/pi/:iterations", PiNumber)
	}

	r.Run()
}

func GetTasks(c *gin.Context) {

	var tasks []models.Task
	db := db.GetDB()
	db.Find(&tasks)
	c.JSON(200, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	var db = db.GetDB()

	if err := c.BindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&task)
	c.JSON(http.StatusOK, &task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&task)
	db.Save(&task)
	c.JSON(http.StatusOK, &task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&task)
}

// curl -i -X GET -H "Content-Type: application/json" http://192.168.99.100:30081/api/v1/tasks/
// curl -i -X POST -H "Content-Type: application/json" -d "{\"title\": \"test\",\"created_at\": \"2017-11-13T23:03:28-08:00\", \"completed\": false}" http://192.168.99.100:30081/api/v1/tasks/
// curl -i -X DELETE -H "Content-Type: application/json" http://192.168.99.100:30081/api/v1/tasks/ID
// curl -i -X PUT -H "Content-Type: application/json" -d "{\"title\": \"test name changed\", \"completed\": true}" http://192.168.99.100:30081/api/v1/tasks/ID

func Kill(c *gin.Context) {
	fmt.Println("EXITING BRUTE FORCE!")
	os.Exit(-1)

	// curl -i -X GET http://localhost:8080/api/v1/kill
}

func PiNumber(c *gin.Context) {

	val := c.Params.ByName("iterations")

	iterations, err := strconv.Atoi(val)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid number of iterations: " + val})
	}
	// POD dies if more than 1000000
	// if iterations > 200000 {
	// 	c.JSON(404, gin.H{"error": "Number too big: " + val})
	// }

	c.JSON(200, gin.H{"success": "Pi is: " + fmt.Sprintf("%.20f", pi(iterations))})
	// curl -i -X GET http://localhost:8080/api/v1/pi/35000
}

func pi(n int) float64 {
	f := 0.0
	for k := 0; k <= n; k++ {
		f += 4 * math.Pow(-1, float64(k)) / (2*float64(k) + 1)
	}
	return f
}

// pi launches n goroutines to compute an
// approximation of pi.
// func pi(n int) float64 {
// 	ch := make(chan float64)
// 	defer close(ch)

// 	for k := 0; k <= n; k++ {
// 		go term(ch, float64(k))
// 	}
// 	f := 0.0
// 	for k := 0; k <= n; k++ {
// 		f += <-ch
// 	}
// 	return f
// }

// func term(ch chan float64, k float64) {
// 	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
// }

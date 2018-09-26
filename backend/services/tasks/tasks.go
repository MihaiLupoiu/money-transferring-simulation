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

	c.JSON(200, gin.H{"success": "Pi is: " + fmt.Sprintf("%.20f", pi(iterations))})
	// curl -i -X GET http://localhost:8080/api/v1/pi/35000
}

// Simple version
func pi(n int) float64 {
	f := 0.0
	for k := 0; k <= n; k++ {
		f += 4 * math.Pow(-1, float64(k)) / (2*float64(k) + 1)
	}
	return f
}

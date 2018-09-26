package main

import (
	"log"
	"net/http"

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

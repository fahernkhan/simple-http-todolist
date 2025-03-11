package main

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Nana's Golang Full Course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	router := gin.Default()

	router.GET("/", helloUser)
	router.GET("/show-tasks", showTask)

	router.Run(":8080")
}

func showTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"task": taskItems,
	})
}

func helloUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello user. Welcome to our Todolist App!",
	})
}

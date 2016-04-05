package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)


func timeoutKill(timeout int) {
	time.Sleep(time.Second * time.Duration(timeout))
	println("Hi")

}

type ContainersJSON struct {
	Container string `json:"container" binding:"required"`
	Timeout string `json:"timeout" binding:"required,numeric"`
}


func createContainersView(c *gin.Context) {
	var json ContainersJSON
	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_params",
		})
		return

	}
	timeout, _ := strconv.Atoi(json.Timeout)
	if timeout <= 0 || timeout > 600 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "timeout shoud be in 0 < timeout <= 600",
		})
		return
	}

	go timeoutKill(timeout)  // To kill containers as async.
	listContainers()

	c.JSON(http.StatusOK, gin.H{
		"postedContainer": json.Container,
	})
}

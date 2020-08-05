package server

import (
	"github.com/GrandTrunkSemaphoreCompany/clex/clacks/encoding"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type WebMessage struct {
	Message     string `form:"message" json:"message" binding:"required"`
	Destination int    `form:"destination" json:"destination" binding:"required,numeric"`
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func queueProcessHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"queue": "empty",
	})
}

func historyHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "queue",
	})
}

func queueTowerHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"queue": "empty",
	})
}

func queuePrintHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"queue": "empty",
	})
}

func sendHandler(c *gin.Context) {
	var wm WebMessage
	if err := c.BindJSON(&wm); err != nil {
		log.Fatal(err)
	}

	m := encoding.Message{
		Body:        wm.Message,
		Destination: wm.Destination,
	}

	c.JSON(200, gin.H{
		"status":   "received",
		"message":  m.Body,
		"received": time.Now().UTC(),
	})
}

func Start() {
	r := gin.Default()

	r.GET("/ping", healthCheckHandler)

	r.POST("/send", sendHandler)

	r.GET("/queue/process", queueProcessHandler)
	r.GET("/queue/print", queuePrintHandler)
	r.GET("/queue/tower/:id", queueTowerHandler)

	r.GET("/history", historyHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

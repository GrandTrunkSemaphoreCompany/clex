package clacks

import (
	"github.com/gin-gonic/gin"
	"log"
)

type WebMessage struct {
	Message     string `form:"message" json:"message" binding:"required"`
	Destination int    `form:"destination" json:"destination" binding:"required,numeric"`
}

func healthCheck(c *gin.Context) {
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
		panic(err)
	}

	m := Message{
		body:        wm.Message,
		destination: wm.Destination,
	}

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": m.String(),
	})

	log.Println(wm.Destination)
}

func Start() {
	r := gin.Default()

	r.GET("/ping", healthCheck)

	r.POST("/send", sendHandler)

	r.GET("/queue/process", queueProcessHandler)
	r.GET("/queue/print", queuePrintHandler)
	r.GET("/queue/tower/:id", queueTowerHandler)

	r.GET("/history", historyHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

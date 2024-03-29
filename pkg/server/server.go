package server

import (
	"fmt"
	"github.com/GrandTrunkSemaphoreCompany/clex/config"
	"github.com/GrandTrunkSemaphoreCompany/clex/pkg/message"
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

	m := message.Message{
		Body:        wm.Message,
		Destination: wm.Destination,
		Created:     time.Now().UTC(),
	}

	c.JSON(200, gin.H{
		"status":   "received",
		"message":  m.Body,
		"received": m.Created,
	})

	//d := encoding.NewDirectory("/tmp/clex", 102)
	//d.Write(m)
}

func Start(c config.Config) {
	printConfigDetails(c)

	//configureSources(c)
	//configureSinks(c)

	r := gin.Default()
	configureRouting(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func configureRouting(r *gin.Engine) {
	r.GET("/ping", healthCheckHandler)

	r.POST("/send", sendHandler)

	r.GET("/queue/process", queueProcessHandler)
	r.GET("/queue/print", queuePrintHandler)
	r.GET("/queue/tower/:id", queueTowerHandler)

	r.GET("/history", historyHandler)

}

func printConfigDetails(c config.Config) {
	fmt.Printf("Server started as %d\n", c.Id)

	fmt.Println("Sinks:")
	for _, v := range c.Sinks {
		fmt.Printf("    %d @ %s\n", v.Id, v.URI)
	}

	fmt.Println("Source:")
	for _, v := range c.Sources {
		fmt.Printf("    %d @ %s\n", v.Id, v.URI)
	}
}

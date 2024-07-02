package main

import (
	"fmt"
	"log"
	"myproject/config"
	"myproject/controllers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Access environment variables
	ginMode := os.Getenv("GIN_MODE")
	fmt.Printf("GIN_MODE set to: %s\n", ginMode)

	config.InitDB()
	r := gin.Default() // creating engine instance - already having Logger and Recorvery middleware
	// gin.SetMode(gin.ReleaseMode) // using code to set up releasae mode - set up environment GIN_MODE

	// define routes for CRUD
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// // define router for /ping
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// // Định nghĩa route cho "/"
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Welcome to the home page!")
	// })

	// // Định nghĩa route cho "/favicon.ico"
	// r.GET("/favicon.ico", func(c *gin.Context) {
	// 	c.Status(204) // Trả về No Content
	// })

	fmt.Println("Server is running on port 8080")
	r.Run(":8080") // Lắng nghe và phục vụ trên 0.0.0.0:8080 : default port -> if we want to use anothe port we use: r.Run(": 8080")
}

// in prodcutiobn environment -> use release mode to optimize the effective and perfomance

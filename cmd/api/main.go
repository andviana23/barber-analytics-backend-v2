package main

import (
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	
	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status": "ok",
			"service": "barber-analytics-backend-v2",
			"version": "2.0.0",
		})
	})
	
	// API v1
	v1 := e.Group("/api/v1")
	v1.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "pong"})
	})
	
	// Start server
	log.Println("🚀 Server starting on :8080")
	log.Fatal(e.Start(":8080"))
}

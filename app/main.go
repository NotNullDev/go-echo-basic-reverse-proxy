package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func main() {
	router := echo.New()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Any("/*", func(c echo.Context) error {
		q := c.Request().URL.String()
		log.Printf("hi! [%s]", q)
		return c.String(200, "port : "+port+" "+" route: "+q)
	})

	err := router.Start(":" + port)
	if err != nil {
		return
	}
}

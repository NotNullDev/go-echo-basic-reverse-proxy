package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/url"
)

func main() {
	router := echo.New()

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	router.AutoTLSManager.Cache = autocert.DirCache("./.www-cache")

	app1Url, _ := url.Parse("http://app1:9001")
	app2Url, _ := url.Parse("http://app2:9002")

	g1 := router.Group("/app1")
	g1.Use(middleware.Proxy(middleware.NewRandomBalancer([]*middleware.ProxyTarget{{
		URL: app1Url,
	}})))

	g2 := router.Group("/app2")
	g2.Use(middleware.Proxy(middleware.NewRandomBalancer([]*middleware.ProxyTarget{{
		URL: app2Url,
	}})))

	router.GET("/*", func(c echo.Context) error {
		return c.String(200, "ok!")
	})

	log.Fatal(router.StartAutoTLS(":443"))
}

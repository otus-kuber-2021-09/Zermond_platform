package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"log"
)

func main() {
	e := echo.New()
	e.Use(ServerHeader)
	e.Use(middleware.Static("/app"))
	log.Panic(e.Start(":8000"))
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		dir, _ := ioutil.ReadDir("/app")
		for _, i := range dir {
			log.Printf("FileName: %s", i.Name())
		}
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

package main

import (
	"insurancesvc/config"
	"insurancesvc/module/mongodb"
	"insurancesvc/route"

	_ "insurancesvc/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Garage Service API
// @version 1.0.0
// @description Documentation of Garage Service API
//
// @host micasvn.ddns.net:9999
// @BasePath /
//
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization

func init() {
	config.Init()
	mongodb.Connect()
}

func main() {
	//Load dotenv for port option
	env := config.GetENV()

	//Echo
	e := echo.New()

	//CORS
	e.Use(middleware.CORS())

	//Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}  -->  method=${method}, uri=${uri}, status=${status}\n",
	}))

	//Swagger
	e.GET("/spec/*", echoSwagger.WrapHandler)

	//Add routers
	route.Bootstrap(e)

	//Open port
	e.Start(env.Port)
}

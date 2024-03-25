package configuration

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (c *configuration) initServer() *configuration {
	fmt.Println("setting up the server...")
	env, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		env = "development"
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	c.portServer = port

	// setup server
	c.Server = echo.New()
	c.Server.Debug = env == "development"
	c.Server.Use(middleware.RemoveTrailingSlash())

	return c
}

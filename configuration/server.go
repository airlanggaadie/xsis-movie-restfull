package configuration

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (c *configuration) initServer() *configuration {
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

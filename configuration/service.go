package configuration

import (
	"fmt"
	handler "xsis/assignment-test/service/delivery/server"
	"xsis/assignment-test/service/repository/postgresql"
	"xsis/assignment-test/service/usecase"
)

func (c *configuration) initService() *configuration {
	fmt.Println("setting up some features...")
	movieRepository := postgresql.NewMovieRepository(c.DB)
	movieUsecase := usecase.NewMovieUsecase(movieRepository)

	handler.NewHandler(c.Server, c.DB, movieUsecase)

	return c
}

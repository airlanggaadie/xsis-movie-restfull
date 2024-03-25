package configuration

import (
	handler "xsis/assignment-test/service/delivery/server"
	"xsis/assignment-test/service/repository/postgresql"
	"xsis/assignment-test/service/usecase"
)

func (c *configuration) initService() *configuration {
	movieRepository := postgresql.NewMovieRepository(c.DB)
	movieUsecase := usecase.NewMovieUsecase(movieRepository)

	handler.NewHandler(c.Server, c.DB, movieUsecase)

	return c
}

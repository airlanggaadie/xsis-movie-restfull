package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"xsis/assignment-test/model"
	"xsis/assignment-test/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	router *echo.Echo
	db     *sql.DB

	movieUsecase service.MovieUsecase
}

func NewHandler(router *echo.Echo, db *sql.DB, movieUsecase service.MovieUsecase) {
	var handler = Handler{
		router:       router,
		db:           db,
		movieUsecase: movieUsecase,
	}

	handler.routes()
}

func (h Handler) healthCheck(c echo.Context) error {
	var dbStatus = "OK"
	err := h.db.PingContext(c.Request().Context())
	if err != nil {
		log.Printf("[handler][HealthCheck] error: %v", err)
		dbStatus = err.Error()
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": dbStatus,
	})
}

func (h Handler) listMovie(c echo.Context) error {
	queryPage := c.QueryParam("page")
	if queryPage == "" {
		queryPage = "0"
	}

	queryLimit := c.QueryParam("limit")
	if queryLimit == "" {
		queryLimit = "0"
	}

	page, err := strconv.Atoi(queryPage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid query parameter page",
		})
	}

	limit, err := strconv.Atoi(queryLimit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid query parameter limit",
		})
	}

	list, err := h.movieUsecase.GetMoviesPaginate(c.Request().Context(), page, limit)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, list)
}

func (h Handler) detailMovie(c echo.Context) error {
	// TODO: validate request and adjust param
	movie, err := h.movieUsecase.GetMovie(uuid.New())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, movie)
}

func (h Handler) addNewMovie(c echo.Context) error {
	// validate request and adjust param
	var body model.AddNewMovieRequest
	err := c.Bind(&body)
	if err != nil {
		return err
	}

	err = body.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	movie, err := h.movieUsecase.AddNewMovie(c.Request().Context(), body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, movie)
}

func (h Handler) updateMovie(c echo.Context) error {
	// TODO: validate request and adjust param
	movie, err := h.movieUsecase.UpdateMovie(uuid.New(), model.Movie{})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, movie)
}

func (h Handler) deleteMovie(c echo.Context) error {
	// TODO: validate request and adjust param
	err := h.movieUsecase.DeleteMovie(uuid.New())
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

package handlers

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func getID(c echo.Context) (int64, error) {
	id, err := strconv.ParseInt(c.QueryParam("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getPaginationParams(c echo.Context) (PaginationParams, error) {
	var page int
	var err error
	if c.QueryParam("page") == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return PaginationParams{}, err
		}
	}

	var limit int
	if c.QueryParam("limit") == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			return PaginationParams{}, err
		}
	}

	return PaginationParams{
		Page:  page,
		Limit: limit,
	}, nil
}

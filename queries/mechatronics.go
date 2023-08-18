package queries

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"koreatech-board-api/db"
	"koreatech-board-api/model"
	"math"
	"net/http"
	"strconv"
)

// @Summary		Get article list
// @Description	Get mechatronics article list
// @Tags			mechatronics
// @Accept			json
// @Produce		json
// @Param			board			path		string	true	"name of the board"
// @Param			page			query		integer	false	"page of board"
// @Param			num_of_items	query		integer	false	"items per page"
// @Success		200				{object}	model.APIData
// @Failure		404
// @Router			/mechatronics/{board} [get]
func SelectMechaQuery(c echo.Context) error {
	boardRaw := c.Param("board")

	var board = ""

	switch boardRaw {
	case "notice":
		board = "235"
	case "lecture":
		board = "236"
	case "bachelor":
		board = "237"
	case "job":
		board = "238"
	case "free":
		board = "244"
	default:
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": fmt.Sprintf("Board \"%s\" not found!", boardRaw),
		})
	}

	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	numOfItems, noiErr := strconv.Atoi(c.QueryParam("num_of_items"))

	if pageErr != nil {
		page = 1
	}

	if noiErr != nil {
		numOfItems = 20
	}

	var results []model.Board
	var count []int64

	listArgs := map[string]interface{}{"board": board, "offset": int64((page - 1) * numOfItems), "num_of_items": int64(numOfItems)}

	var listQuery = db.Pool.Query(c.Request().Context(),
		`SELECT mechatronics 
		{ id, num, title, writer, write_date, read_count }
		FILTER .board = <str>$board order by contains(.num, '공지') DESC
		THEN .write_date DESC
		THEN .num DESC OFFSET <int64>$offset limit <int64>$num_of_items`,
		&results,
		listArgs,
	)

	countArgs := map[string]interface{}{"board": board}

	var countQuery = db.Pool.Query(c.Request().Context(),
		`SELECT count(mechatronics filter .board=<str>$board)`,
		&count,
		countArgs,
	)

	if listQuery != nil || countQuery != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Query error!",
		})
	}

	apiData := model.APIData{
		LastPage: int(math.Ceil(float64(count[0]) / float64(numOfItems))),
		Posts:    results,
	}

	return c.JSON(http.StatusOK, apiData)
}

// @Summary		Get article
// @Description	Get mechatronics article by UUID
// @Tags			mechatronics
// @Accept			json
// @Produce		json
// @Param			uuid	query		string	true	"uuid of article"
// @Success		200		{object}	model.Article
// @Failure		404
// @Router			/article/mechatronics [get]
func MechaArticleQuery(c echo.Context) error {
	var results []model.Article

	var articleQuery = db.Pool.Query(c.Request().Context(),
		`SELECT mechatronics
		{ id, title, writer, write_date, article_url, content, files: {file_name, file_url} }
		FILTER .id = <uuid><str>$0`,
		&results,
		c.QueryParam("uuid"),
	)

	if articleQuery != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Query error!",
		})
	}

	return c.JSON(http.StatusOK, results[0])
}

// @Summary		Search article by title
// @Description	Search article from specific board by title
// @Tags			mechatronics
// @Accept			json
// @Produce		json
// @Param			board			path		string	true	"name of the board"
// @Param			title	query		string	true	"title"
// @Param			page			query		integer	false	"page of board"
// @Param			num_of_items	query		integer	false	"items per page"
// @Success		200		{object}	model.Article
// @Failure		404
// @Router			/mechatronics/{board}/search/title [get]
func MechatronicsSearchWithTitleQuery(c echo.Context) error {
	boardRaw := c.Param("board")

	var board = ""

	switch boardRaw {
	case "notice":
		board = "235"
	case "lecture":
		board = "236"
	case "bachelor":
		board = "237"
	case "job":
		board = "238"
	case "free":
		board = "244"
	default:
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": fmt.Sprintf("Board \"%s\" not found!", boardRaw),
		})
	}
	title := "%" + c.QueryParam("title") + "%"

	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	numOfItems, noiErr := strconv.Atoi(c.QueryParam("num_of_items"))

	if pageErr != nil {
		page = 1
	}

	if noiErr != nil {
		numOfItems = 20
	}

	var results []model.Board
	var count []int64

	listArgs := map[string]interface{}{"board": board, "title": title, "offset": int64((page - 1) * numOfItems), "num_of_items": int64(numOfItems)}

	var listQuery = db.Pool.Query(c.Request().Context(),
		`SELECT mechatronics 
		{ id, num, title, writer, write_date, read_count }
		FILTER .board = <str>$board and .title ilike <str>$title order by contains(.num, '공지') DESC
		THEN .write_date DESC
		THEN .num DESC OFFSET <int64>$offset limit <int64>$num_of_items`,
		&results,
		listArgs,
	)

	countArgs := map[string]interface{}{"board": board, "title": title}

	var countQuery = db.Pool.Query(c.Request().Context(),
		`SELECT count(mechatronics filter .board=<str>$board and .title ilike <str>$title)`,
		&count,
		countArgs,
	)

	if listQuery != nil || countQuery != nil {
		fmt.Println(listQuery)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Query error!",
		})
	}

	apiData := model.APIData{
		LastPage: int(math.Ceil(float64(count[0]) / float64(numOfItems))),
		Posts:    results,
	}

	return c.JSON(http.StatusOK, apiData)
}

package queries

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/db"
	"koreatech-board-api/cmd/model"
	"math"
	"net/http"
	"strconv"
)

// @Summary		Get article list
// @Description	Get dorm article list
// @Tags			dorm
// @Accept			json
// @Produce		json
// @Param			board			path		string	true	"name of the board"
// @Param			page			query		integer	false	"page of board"
// @Param			num_of_items	query		integer	false	"items per page"
// @Success		200				{object}	model.APIData
// @Failure		400
// @Failure		404
// @Router			/dorm/{board} [get]
func SelectDormQuery(c echo.Context) error {
	boardRaw := c.Param("board")

	apiError := ""
	status := http.StatusOK

	var board = ""

	switch boardRaw {
	case "notice":
		board = "notice"
	case "free":
		board = "bulletin"
	default:
		status = http.StatusNotFound
		apiError = fmt.Sprintf("Board \"%s\" not found!", boardRaw)
	}

	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	numOfItems, noiErr := strconv.Atoi(c.QueryParam("num_of_items"))

	if pageErr != nil {
		page = 1
	}

	if noiErr != nil {
		numOfItems = 20
	}

	results := []model.Board{}
	var count []int64

	listArgs := map[string]interface{}{"board": board, "offset": int64((page - 1) * numOfItems), "num_of_items": int64(numOfItems)}

	var listQuery = db.Pool.Query(c.Request().Context(),
		`SELECT dorm 
		{ id, num, title, writer, write_date, read_count }
		FILTER .board = <str>$board order by contains(.num, '공지') DESC
		THEN .write_date DESC
		THEN .num DESC OFFSET <int64>$offset limit <int64>$num_of_items`,
		&results,
		listArgs,
	)

	countArgs := map[string]interface{}{"board": board}

	var countQuery = db.Pool.Query(c.Request().Context(),
		`SELECT count(dorm filter .board=<str>$board)`,
		&count,
		countArgs,
	)

	if listQuery != nil || countQuery != nil {
		status = http.StatusBadRequest
		apiError = "Query error!"
	}

	apiData := model.APIData{
		StatusCode: status,
		LastPage:   int(math.Ceil(float64(count[0]) / float64(numOfItems))),
		Error:      apiError,
		Posts:      results,
	}

	return c.JSON(status, apiData)
}

// @Summary		Get article
// @Description	Get dorm article by UUID
// @Tags			dorm
// @Accept			json
// @Produce		json
// @Param			uuid	query		string	true	"uuid of article"
// @Success		200		{object}	model.ApiArticle
// @Failure		400
// @Failure		404
// @Router			/article/dorm [get]
func DormArticleQuery(c echo.Context) error {
	var results []model.Article

	apiError := ""
	status := http.StatusOK

	var articleQuery = db.Pool.Query(c.Request().Context(),
		`SELECT dorm
		{ id, title, writer, write_date, article_url, content, files: {file_name, file_url} }
		FILTER .id = <uuid><str>$0`,
		&results,
		c.QueryParam("uuid"),
	)

	if articleQuery != nil {
		status = http.StatusBadRequest
		apiError = "Query error!"
	}

	if len(results) == 0 {
		article := model.ApiArticle{
			StatusCode: status,
			Error:      apiError,
		}
		return c.JSON(status, article)
	} else {
		article := model.ApiArticle{
			StatusCode: status,
			Error:      apiError,
			Id:         results[0].Id,
			Title:      results[0].Title,
			Writer:     results[0].Writer,
			WriteDate:  results[0].WriteDate,
			ArticleUrl: results[0].ArticleUrl,
			Content:    results[0].Content,
			Files:      results[0].Files,
		}
		return c.JSON(status, article)
	}
}

// @Summary		Search article by title
// @Description	Search article from specific board by title
// @Tags			dorm
// @Accept			json
// @Produce		json
// @Param			board			path		string	true	"name of the board"
// @Param			title	query		string	true	"title"
// @Param			page			query		integer	false	"page of board"
// @Param			num_of_items	query		integer	false	"items per page"
// @Success		200		{object}	model.ApiArticle
// @Failure		400
// @Failure		404
// @Router			/dorm/{board}/search/title [get]
func DormSearchWithTitleQuery(c echo.Context) error {
	boardRaw := c.Param("board")

	apiError := ""
	status := http.StatusOK

	var board = ""

	switch boardRaw {
	case "notice":
		board = "notice"
	case "free":
		board = "bulletin"
	default:
		status = http.StatusNotFound
		apiError = fmt.Sprintf("Board \"%s\" not found!", boardRaw)
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

	results := []model.Board{}
	var count []int64

	listArgs := map[string]interface{}{"board": board, "title": title, "offset": int64((page - 1) * numOfItems), "num_of_items": int64(numOfItems)}

	var listQuery = db.Pool.Query(c.Request().Context(),
		`SELECT dorm 
		{ id, num, title, writer, write_date, read_count }
		FILTER .board = <str>$board and .title ilike <str>$title order by contains(.num, '공지') DESC
		THEN .write_date DESC
		THEN .num DESC OFFSET <int64>$offset limit <int64>$num_of_items`,
		&results,
		listArgs,
	)

	countArgs := map[string]interface{}{"board": board, "title": title}

	var countQuery = db.Pool.Query(c.Request().Context(),
		`SELECT count(dorm filter .board=<str>$board and .title ilike <str>$title)`,
		&count,
		countArgs,
	)

	if listQuery != nil || countQuery != nil {
		status = http.StatusBadRequest
		apiError = "Query error!"
	}

	apiData := model.APIData{
		StatusCode: status,
		LastPage:   int(math.Ceil(float64(count[0]) / float64(numOfItems))),
		Error:      apiError,
		Posts:      results,
	}

	return c.JSON(status, apiData)
}

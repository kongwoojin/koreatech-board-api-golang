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

func SelectSchoolQuery(c echo.Context) error {
	boardRaw := c.Param("board")

	var board = ""

	switch boardRaw {
	case "notice":
		board = "list"
	case "scholar":
		board = "scholarList"
	case "bachelor":
		board = "bachelorList"
	case "covid19":
		board = "boardList8"
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
		`SELECT school 
		{ id, num, title, writer, write_date, read_count } 
		FILTER .board=<str>$board order by .is_importance DESC 
		THEN .write_date DESC
		THEN .num desc offset <int64>$offset limit <int64>$num_of_items`,
		&results,
		listArgs,
	)

	countArgs := map[string]interface{}{"board": board}

	var countQuery = db.Pool.Query(c.Request().Context(),
		`SELECT count(school filter .board=<str>$board)`,
		&count,
		countArgs,
	)

	if listQuery != nil || countQuery != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Query error!",
		})
	}

	apiData := model.APIData{
		LastPage: int(math.Ceil(float64(int(count[0]) / numOfItems))),
		Posts:    results,
	}

	return c.JSON(http.StatusOK, apiData)
}

func SchoolArticleQuery(c echo.Context) error {
	var results []model.Article

	var articleQuery = db.Pool.Query(c.Request().Context(),
		`SELECT school
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

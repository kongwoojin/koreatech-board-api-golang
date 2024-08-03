package queries

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/db"
	"koreatech-board-api/cmd/enums"
	"koreatech-board-api/cmd/model"
	"math"
	"net/http"
	"strconv"
)

// @Summary		Search article by title
// @Description	Search article from specific board by title
// @Tags			Notice
// @Accept			json
// @Produce		json
// @Param   department  path     string     true  "name of the department"       Enums(arch, cse, dorm, mse, ace, ide, ite, mechanical, mechatronics, school, sim)
// @Param			board			path		string	true	"name of the board" Enums(notice, free, job, pds, lecture, bachelor, scholar)
// @Param			title	query		string	true	"title"
// @Param			page			query		integer	false	"page of board"
// @Param			num_of_items	query		integer	false	"items per page"
// @Success		200		{object}	model.ApiArticle
// @Failure		400
// @Failure		404
// @Router			/{department}/{board}/search/title [get]
func SearchWithTitle(c echo.Context) error {
	departmentRaw := c.Param("department")
	boardRaw := c.Param("board")

	apiError := ""
	status := http.StatusOK

	department, ok := enums.ParseDepartment(departmentRaw)

	if !ok {
		status = http.StatusNotFound
		apiError = fmt.Sprintf("Department \"%s\" not found!", departmentRaw)
	}

	board, ok := enums.ParseBoard(boardRaw)

	if !ok {
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
	lastPage := 1

	if department != enums.UNKNOWN_DEPARTMENT && board != enums.UNKNOWN_BOARD {
		listArgs := map[string]interface{}{"department": department.String(), "board": board.String(), "title": title, "offset": int64((page - 1) * numOfItems), "num_of_items": int64(numOfItems)}

		var listQuery = db.Pool.Query(c.Request().Context(),
			`SELECT notice 
		{ id, num, title, writer, write_date, read_count, is_new := .init_crawled_time = .update_crawled_time, is_notice }
		FILTER .department=<Department><str>$department AND .board=<Board><str>$board
		AND .title ilike <str>$title order by .is_notice DESC
		THEN .write_date DESC
		THEN .num DESC OFFSET <int64>$offset limit <int64>$num_of_items`,
			&results,
			listArgs,
		)

		countArgs := map[string]interface{}{"department": department.String(), "board": board.String(), "title": title}

		var countQuery = db.Pool.Query(c.Request().Context(),
			`SELECT count(notice filter .department=<Department><str>$department AND .board=<Board><str>$board AND .title ilike <str>$title)`,
			&count,
			countArgs,
		)

		if listQuery != nil || countQuery != nil {
			status = http.StatusBadRequest
			apiError = "Query error!"
		}

		lastPage = int(math.Ceil(float64(count[0]) / float64(numOfItems)))
	}

	apiData := model.APIData{
		StatusCode: status,
		LastPage:   lastPage,
		Error:      apiError,
		Posts:      results,
	}

	return c.JSON(status, apiData)
}

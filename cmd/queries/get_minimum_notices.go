package queries

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/db"
	"koreatech-board-api/cmd/enums"
	"koreatech-board-api/cmd/model"
	"net/http"
)

// @Summary		Get minumum notice list
// @Description	Get minumum notice list for board widget, only 5 new notices
// @Tags			Notice
// @Accept			json
// @Produce		json
// @Param   department  path     string     true  "name of the department"       Enums(arch, cse, dorm, mse, ace, ide, ite, mechanical, mechatronics, school, sim)
// @Param			board			path		string	true	"name of the board" Enums(notice, free, job, pds, lecture, bachelor, scholar)
// @Success		200				{object}	model.APIData
// @Failure		400
// @Failure		404
// @Router			/{department}/{board}/widget [get]
func GetMinimumNotices(c echo.Context) error {
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

	var results []model.Board = nil
	var count []int64

	if department != enums.UNKNOWN_DEPARTMENT && board != enums.UNKNOWN_BOARD {
		listArgs := map[string]interface{}{"department": department.String(), "board": board.String()}

		var listQuery = db.Pool.Query(c.Request().Context(),
			`SELECT notice 
		{ id, num, title, writer, write_date, read_count, is_new := .init_crawled_time = .update_crawled_time, is_notice }
		FILTER .department=<Department><str>$department AND .board=<Board><str>$board order by .write_date DESC
		THEN .num desc offset 0 limit 5`,
			&results,
			listArgs,
		)

		countArgs := map[string]interface{}{"department": department.String(), "board": board.String()}

		var countQuery = db.Pool.Query(c.Request().Context(),
			`SELECT count(notice filter .department=<Department><str>$department AND .board=<Board><str>$board)`,
			&count,
			countArgs,
		)

		if listQuery != nil || countQuery != nil {
			status = http.StatusBadRequest
			apiError = "Query error!"
		}
	}

	for result := range results {
		results[result].Num = 5 - int64(result)
	}

	apiData := model.APIData{
		StatusCode: status,
		LastPage:   1,
		Error:      apiError,
		Posts:      results,
	}

	return c.JSON(status, apiData)
}
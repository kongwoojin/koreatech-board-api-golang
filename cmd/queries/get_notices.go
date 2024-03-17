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

// @Summary		Get notice list
// @Description	Get notice list
// @Tags			Notice
// @Accept			json
// @Produce		json
// @Param   department  path     string     true  "name of the department"       Enums(arch, cse, dorm, mse, ace, ide, ite, mechanical, mechatronics, school, sim)
// @Param			board			path		string	true	"name of the board" Enums(notice, free, job, pds, lecture, bachelor, scholar)
// @Param			page			query		integer	false	"page of board"
// @Param			num_of_items	query		integer	false	"items per page"
// @Success		200				{object}	model.APIData
// @Failure		400
// @Failure		404
// @Router			/{department}/{board} [get]
func GetNotices(c echo.Context) error {
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

	fmt.Printf("Board: %v\n", board)

	if !ok {
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

	var results []model.Board = nil
	var count []int64
	lastPage := 1

	if department != enums.UNKNOWN_DEPARTMENT && board != enums.UNKNOWN_BOARD {
		listArgs := map[string]interface{}{"department": department.String(), "board": board.String(), "offset": int64((page - 1) * numOfItems), "num_of_items": int64(numOfItems)}

		var listQuery = db.Pool.Query(c.Request().Context(),
			`SELECT notice 
		{ id, num, title, writer, write_date, read_count } 
		FILTER .department=<Department><str>$department AND .board=<Board><str>$board order by .is_notice DESC 
		THEN .write_date DESC
		THEN .num desc offset <int64>$offset limit <int64>$num_of_items`,
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

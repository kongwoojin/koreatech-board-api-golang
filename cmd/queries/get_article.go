package queries

import (
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/db"
	"koreatech-board-api/cmd/model"
	"net/http"
)

// @Summary		Get article
// @Description	Get article by UUID
// @Tags			Notice
// @Accept			json
// @Produce		json
// @Param			uuid	query		string	true	"uuid of article"
// @Success		200		{object}	model.ApiArticle
// @Failure		400
// @Failure		404
// @Router			/article [get]
func GetArticle(c echo.Context) error {
	var results []model.Article

	apiError := ""
	status := http.StatusOK

	var articleQuery = db.Pool.Query(c.Request().Context(),
		`SELECT notice
		{ id, num, title, writer, write_date, article_url, content, is_notice, files: {file_name, file_url} }
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
			Num:        results[0].Num,
			Id:         results[0].Id,
			Title:      results[0].Title,
			Writer:     results[0].Writer,
			WriteDate:  results[0].WriteDate,
			ArticleUrl: results[0].ArticleUrl,
			Content:    results[0].Content,
			IsNotice:   results[0].IsNotice,
			Files:      results[0].Files,
		}
		return c.JSON(status, article)
	}
}

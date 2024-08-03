package routes

import (
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/queries"
)

func APIRouter(e *echo.Echo) {
	g := e.Group("/v3")

	g.GET("/article", queries.GetArticle)
	g.GET("/:department/:board", queries.GetNotices)
	g.GET("/:department/:board/widget", queries.GetMinimumNotices)
	g.GET("/:department/:board/search/title", queries.SearchWithTitle)
}

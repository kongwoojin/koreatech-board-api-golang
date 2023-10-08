package routes

import (
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/queries"
)

func APIRouter(e *echo.Echo) {
	g := e.Group("/v3")

	g.GET("/article/arch", queries.ArchArticleQuery)
	g.GET("/arch/:board", queries.SelectArchQuery)
	g.GET("/arch/:board/search/title", queries.ArchSearchWithTitleQuery)

	g.GET("/article/cse", queries.CseArticleQuery)
	g.GET("/cse/:board", queries.SelectCseQuery)
	g.GET("/cse/:board/search/title", queries.CseSearchWithTitleQuery)

	g.GET("/article/dorm", queries.DormArticleQuery)
	g.GET("/dorm/:board", queries.SelectDormQuery)
	g.GET("/dorm/:board/search/title", queries.DormSearchWithTitleQuery)

	g.GET("/article/emc", queries.EmcArticleQuery)
	g.GET("/emc/:board", queries.SelectEmcQuery)
	g.GET("/emc/:board/search/title", queries.EmcSearchWithTitleQuery)

	g.GET("/article/ide", queries.IdeArticleQuery)
	g.GET("/ide/:board", queries.SelectIdeQuery)
	g.GET("/ide/:board/search/title", queries.IdeSearchWithTitleQuery)

	g.GET("/article/ite", queries.IteArticleQuery)
	g.GET("/ite/:board", queries.SelectIteQuery)
	g.GET("/ite/:board/search/title", queries.IteSearchWithTitleQuery)

	g.GET("/article/mechanical", queries.MechanicalArticleQuery)
	g.GET("/mechanical/:board", queries.SelectMechanicalQuery)
	g.GET("/mechanical/:board/search/title", queries.MechanicalSearchWithTitleQuery)

	g.GET("/article/mechatronics", queries.MechaArticleQuery)
	g.GET("/mechatronics/:board", queries.SelectMechaQuery)
	g.GET("/mechatronics/:board/search/title", queries.MechatronicsSearchWithTitleQuery)

	g.GET("/article/school", queries.SchoolArticleQuery)
	g.GET("/school/:board", queries.SelectSchoolQuery)
	g.GET("/school/:board/search/title", queries.SchoolSearchWithTitleQuery)

	g.GET("/article/sim", queries.SimArticleQuery)
	g.GET("/sim/:board", queries.SelectSimQuery)
	g.GET("/sim/:board/search/title", queries.SimSearchWithTitleQuery)
}

package routes

import (
	"github.com/labstack/echo/v4"
	"koreatech-board-api/cmd/queries"
)

func APIRouter(e *echo.Echo) {
	e.GET("/v3/article/arch", queries.ArchArticleQuery)
	e.GET("/v3/arch/:board", queries.SelectArchQuery)
	e.GET("/v3/arch/:board/search/title", queries.ArchSearchWithTitleQuery)

	e.GET("/v3/article/cse", queries.CseArticleQuery)
	e.GET("/v3/cse/:board", queries.SelectCseQuery)
	e.GET("/v3/cse/:board/search/title", queries.CseSearchWithTitleQuery)

	e.GET("/v3/article/dorm", queries.DormArticleQuery)
	e.GET("/v3/dorm/:board", queries.SelectDormQuery)
	e.GET("/v3/dorm/:board/search/title", queries.DormSearchWithTitleQuery)

	e.GET("/v3/article/emc", queries.EmcArticleQuery)
	e.GET("/v3/emc/:board", queries.SelectEmcQuery)
	e.GET("/v3/emc/:board/search/title", queries.EmcSearchWithTitleQuery)

	e.GET("/v3/article/ide", queries.IdeArticleQuery)
	e.GET("/v3/ide/:board", queries.SelectIdeQuery)
	e.GET("/v3/ide/:board/search/title", queries.IdeSearchWithTitleQuery)

	e.GET("/v3/article/ite", queries.IteArticleQuery)
	e.GET("/v3/ite/:board", queries.SelectIteQuery)
	e.GET("/v3/ite/:board/search/title", queries.IteSearchWithTitleQuery)

	e.GET("/v3/article/mechanical", queries.MechanicalArticleQuery)
	e.GET("/v3/mechanical/:board", queries.SelectMechanicalQuery)
	e.GET("/v3/mechanical/:board/search/title", queries.MechanicalSearchWithTitleQuery)

	e.GET("/v3/article/mechatronics", queries.MechaArticleQuery)
	e.GET("/v3/mechatronics/:board", queries.SelectMechaQuery)
	e.GET("/v3/mechatronics/:board/search/title", queries.MechatronicsSearchWithTitleQuery)

	e.GET("/v3/article/school", queries.SchoolArticleQuery)
	e.GET("/v3/school/:board", queries.SelectSchoolQuery)
	e.GET("/v3/school/:board/search/title", queries.SchoolSearchWithTitleQuery)

	e.GET("/v3/article/sim", queries.SimArticleQuery)
	e.GET("/v3/sim/:board", queries.SelectSimQuery)
	e.GET("/v3/sim/:board/search/title", queries.SimSearchWithTitleQuery)
}

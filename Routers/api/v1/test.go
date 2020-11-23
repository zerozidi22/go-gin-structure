package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	// appG := app.Gin{C: c}
	name := c.Query("name")
	fmt.Print(name)
	fmt.Print("getTags Called \n")
	// state := -1
	// if arg := c.Query("state"); arg != "" {
	// 	state = com.StrTo(arg).MustInt()
	// }

	// tagService := tag_service.Tag{
	// 	Name:     name,
	// 	State:    state,
	// 	PageNum:  util.GetPage(c),
	// 	PageSize: setting.AppSetting.PageSize,
	// }
	// tags, err := tagService.GetAll()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
	// 	return
	// }

	// count, err := tagService.Count()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
	// 	return
	// }

	// appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	// 	"lists": tags,
	// 	"total": count,
	// })
}

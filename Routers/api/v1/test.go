package v1

import (
	"net/http"

	"go-gin-structure/pkg/app"

	"go-gin-structure/pkg/enum"

	"github.com/gin-gonic/gin"
)

// [get] using param
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")

	appG.Response(http.StatusOK, enum.SUCCESS, map[string]interface{}{
		"name":  name,
		"total": "count",
	})
}

type AddTagForm struct {
	Name      string `json:"name" form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `json:"created_by" form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `json:"state" form:"state" valid:"Range(0,1)"`
}

// [POST] using body
/*
{
    "name":     "HEESU",
	"created_by": "HAPPY_WATER",
	"state":     0
}
*/
func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTagForm
	)

	appG.Response(http.StatusOK, enum.SUCCESS, map[string]interface{}{
		"name":      form.Name,
		"createdBy": form.CreatedBy,
		"state":     form.State,
	})
}

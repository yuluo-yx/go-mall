package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"go-mall/service"
	"net/http"
)

func ListCategories(c *gin.Context) {
	// 业务层对象
	var listCategoriesService service.CategoriesService

	if err := c.ShouldBind(&listCategoriesService); err == nil {
		res := listCategoriesService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("Get carousels err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

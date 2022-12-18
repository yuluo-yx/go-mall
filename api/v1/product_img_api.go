package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"go-mall/service"
	"net/http"
)

// ListProductImg 查询商品的图片列表
func ListProductImg(c *gin.Context) {
	var listProductService service.ProductImgService
	if err := c.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("Create Product err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

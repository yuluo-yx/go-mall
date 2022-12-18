package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"go-mall/service"
	"net/http"
)

func CreateCarts(c *gin.Context) {
	var createCartsService service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCartsService); err == nil {
		res := createCartsService.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func GetCarts(c *gin.Context) {
	var getCartsService service.CartService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getCartsService); err == nil {
		res := getCartsService.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func UpdateCarts(c *gin.Context) {
	var updateCartsService service.CartService
	if err := c.ShouldBind(&updateCartsService); err == nil {
		res := updateCartsService.Update(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func DeleteCarts(c *gin.Context) {
	var deleteCartsService service.CartService
	if err := c.ShouldBind(&deleteCartsService); err == nil {
		res := deleteCartsService.Delete(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

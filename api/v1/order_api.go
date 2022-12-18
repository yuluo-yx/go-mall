package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"go-mall/service"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var createOrderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func ShowOrder(c *gin.Context) {
	var showOrderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func ListOrders(c *gin.Context) {
	listOrdersService := service.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listOrdersService); err == nil {
		res := listOrdersService.List(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, serializer.ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteOrder(c *gin.Context) {
	var deleteOrderService service.OrderService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

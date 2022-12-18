package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"go-mall/service"
	"net/http"
)

func OrderPay(c *gin.Context) {
	var paymentService service.OrderPaymentService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&paymentService); err == nil {
		res := paymentService.PayDown(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

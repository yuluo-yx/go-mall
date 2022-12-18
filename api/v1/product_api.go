package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"go-mall/service"
	"net/http"
)

// Product 创建商品
func Product(c *gin.Context) {
	var createProductService service.ProductService

	form, _ := c.MultipartForm()
	files := form.File["file"]
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claims.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("Create Product err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

// ListProducts 查找商品
func ListProducts(c *gin.Context) {
	var productsService service.ProductService

	if err := c.ShouldBind(&productsService); err == nil {
		res := productsService.Products(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("Get Product err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

// SearchProduct 搜索商品信息
func SearchProduct(c *gin.Context) {
	var getProductService service.ProductService
	if err := c.ShouldBind(&getProductService); err == nil {
		res := getProductService.SearchProduct(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("Create Product err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

func ShowProduct(c *gin.Context) {
	var showProductService service.ProductService
	if err := c.ShouldBind(&showProductService); err == nil {
		res := showProductService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("Create Product err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}
}

package v1

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"go-mall/service"
	"net/http"
	"strconv"
)

func ShowFavorites(c *gin.Context) {
	var showFavorites service.FavoriteService

	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showFavorites); err == nil {
		res := showFavorites.List(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("ShowFavorites err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}

}

func DeleteFavorites(c *gin.Context) {
	var deleteFavorites service.FavoriteService
	favoriteId, _ := strconv.Atoi(c.Param("id"))
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFavorites); err == nil {
		res := deleteFavorites.Delete(c.Request.Context(), claims.ID, uint(favoriteId)) // id是收藏夹的id，需要移除那个收藏夹
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("DeleteFavorites err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}

}

func CreateFavorites(c *gin.Context) {
	var createFavorites service.FavoriteService

	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createFavorites); err == nil {
		res := createFavorites.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Info("CreateFavorites err: ", err)
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
	}

}

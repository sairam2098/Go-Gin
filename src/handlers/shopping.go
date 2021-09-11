package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gogin/src/models"
	"gogin/src/services"
)

var ShoppingService services.ShoppingService = services.ShoppingService{}

func GetItems(ctx *gin.Context) {
	items := ShoppingService.GetItems()
	ctx.JSON(http.StatusOK, items)
}

func GetItemsByCategory(ctx *gin.Context) {
	var params models.CategoryStruct
	if ctx.BindQuery(&params) != nil {
		ctx.JSON(http.StatusBadRequest, "Please check the query params")
		return
	}

	items, err := ShoppingService.GetItemsByCategory(params.Category)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Not Found")
	}
	ctx.JSON(http.StatusOK, items)
}

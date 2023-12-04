package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Opening",
	})
}
func ShowOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Opening",
	})
}
func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Opening",
	})
}
func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Opening",
	})
}
func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Opening",
	})
}

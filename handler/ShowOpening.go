package handler

import (
	"net/http"

	"github.com/gabferr/minhas_financas/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if error := db.First(&opening, id).Error; error != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSucess(ctx, "show-opening", opening)
}

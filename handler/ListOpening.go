package handler

import (
	"net/http"

	"github.com/gabferr/minhas_financas/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if error := db.Find(&openings).Error; error != nil {
		sendError(ctx, http.StatusBadRequest, "error list openings")
		return
	}
	sendSucess(ctx, "List Openings", openings)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/gabferr/minhas_financas/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if error := db.First(&opening, id).Error; error != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	if error := db.Delete(&opening).Error; error != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting with id: %s", id))
		return
	}
	sendSucess(ctx, "delete-opening", opening)
}

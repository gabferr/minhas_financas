package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if error := request.Validate(); error != nil {
		logger.Errorf("validation error: %v", error.Error())
		sendError(ctx, http.StatusBadRequest, error.Error())
		return
	}

	if error := db.Create(&request).Error; error != nil {
		logger.Errorf("error creating opening: %v", error.Error())
		return
	}

}

package handler

import (
	"net/http"

	"github.com/gabferr/minhas_financas/schemas"
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

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if error := db.Create(&opening).Error; error != nil {
		logger.Errorf("error creating opening: %v", error.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSucess(ctx, "create-opening ", opening)
}

package handler

import (
	"net/http"

	"github.com/gabferr/minhas_financas/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)

	if error := request.Validate(); error != nil {
		logger.Errorf("validation error: %v", error.Error())
		sendError(ctx, http.StatusBadRequest, error.Error())
		return
	}

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
	//Update
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if error := db.Save(&opening).Error; error != nil {
		logger.Errorf("error update opening: %v", error.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSucess(ctx, "update-opening", opening)
}

package handler

import (
	"blue_common/model/requests"
	users_service "blue_common/service/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var req requests.UserRegisterationRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := users_service.UserRegisterService(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, "success")
	}
}

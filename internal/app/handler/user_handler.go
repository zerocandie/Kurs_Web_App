package handler

import (
	"WebApp/internal/app/ds"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(
	ctx *gin.Context,
) {

	var user ds.User

	if err := ctx.ShouldBindJSON(
		&user,
	); err != nil {

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	err := h.Repository.
		CreateUser(&user)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		user,
	)
}

func (h *Handler) Login(
	ctx *gin.Context,
) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "auth stub",
		},
	)
}

func (h *Handler) Logout(
	ctx *gin.Context,
) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "logout stub",
		},
	)
}
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCities(
	ctx *gin.Context,
) {

	cities, err := h.Repository.
		GetAllCities()

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		cities,
	)
}
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCities godoc
// @Summary      Список городов
// @Tags         Cities
// @Produce      json
// @Success      200  {array}  map[string]interface{}  "Список городов"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Router       /cities [get]
func (h *Handler) GetCities(ctx *gin.Context) {
	cities, err := h.Repository.GetAllCities()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cities)
}
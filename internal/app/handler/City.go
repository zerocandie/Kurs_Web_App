package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetAllCity(ctx *gin.Context) {
	cities, err := h.Repository.GetAllCities()
	if err != nil {
		logrus.Error(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.HTML(http.StatusOK, "city.html", gin.H{
		"city": cities,
	})
}

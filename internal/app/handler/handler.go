package handler

import (
	"WebApp/internal/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{Repository: r}

}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/events/create", h.ShowCreateEventForm)
	router.POST("/events/create", h.AddEvent)
	router.GET("/events/", h.GetCitysChat)
	router.GET("/events/:id", h.GetCityChat)
	router.GET("/AllCity/", h.GetAllCity)
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./resources")
}
func (h *Handler) errorHandler(ctx *gin.Context, errorStatusCode int, err error) {
	logrus.Error(err.Error())
	ctx.JSON(errorStatusCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})
}

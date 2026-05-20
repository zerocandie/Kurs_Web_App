package handler

import (
	"WebApp/internal/app/middleware"
	"WebApp/internal/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(
	r *repository.Repository,
) *Handler {

	return &Handler{
		Repository: r,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// === ПУБЛИЧНЫЕ ===
	public := api.Group("/")
	{
		public.GET("/events", h.GetEvents)
		public.GET("/events/:id", h.GetEventByID)
		public.GET("/cities", h.GetCities)
		public.POST("/auth/login", h.Login)      // ← вход
		public.POST("/users/register", h.RegisterUser)
	}

	// === ЗАЩИЩЁННЫЕ (любой авторизованный) ===
	user := api.Group("/")
	user.Use(middleware.AuthMiddleware())
	{
		// Заявки
		user.GET("/votes", h.GetVotes)           // ← только свои заявки (фильтр в репозитории!)
		user.GET("/votes/:id", h.GetVoteByID)
		user.GET("/votes/cart", h.GetCart)
		user.PUT("/votes/:id/form", h.FormVote)
		user.DELETE("/votes/:id", h.DeleteVote)

		// Связи услуги-заявка
		user.POST("/vote-items", h.AddVoteItem)
		user.PUT("/vote-items/:event_id", h.UpdateVoteItem)
		user.DELETE("/vote-items/:event_id", h.DeleteVoteItem)

		// Выход
		user.POST("/auth/logout", h.Logout)
	}

	// === ТОЛЬКО МОДЕРАТОРЫ ===
	moderator := api.Group("/")
	moderator.Use(middleware.AuthMiddleware(), middleware.ModeratorOnly())
	{
		moderator.PUT("/votes/:id/complete", h.CompleteVote) // ← завершение
		moderator.PUT("/votes/:id/reject", h.RejectVote)     // ← отклонение
	}
}

func (h *Handler) RegisterStatic(
	router *gin.Engine,
) {

	router.LoadHTMLGlob(
		"templates/*",
	)

	router.Static(
		"/static",
		"./resources",
	)
}

func (h *Handler) Error(
	ctx *gin.Context,
	code int,
	err error,
) {

	logrus.Error(err)

	ctx.JSON(
		code,
		gin.H{
			"error": err.Error(),
		},
	)
}
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
func (h *Handler) RegisterRoutes(
	router *gin.Engine,
) {

	api := router.Group("/api")
	{
		/* EVENTS */
		api.GET("/events", h.GetEvents)
		api.GET("/events/:id", h.GetEventByID)
		api.POST("/events", h.CreateEvent)

		/* VOTE */
		api.GET("/votes/cart", h.GetCart)

		api.GET("/votes", h.GetVotes)
		api.GET("/votes/:id", h.GetVoteByID)

		api.PUT("/votes/:id/forym", h.FormVote)
		api.PUT("/votes/:id/complete", h.CompleteVote)
		api.PUT("/votes/:id/reject", h.RejectVote)
		api.DELETE("/votes/:id", h.DeleteVote)

		/* M-M */
		api.POST("/vote-items", h.AddVoteItem)
		api.PUT("/vote-items/:event_id", h.UpdateVoteItem)
		api.DELETE("/vote-items/:event_id", h.DeleteVoteItem)

		/* USERS */
		api.POST("/users/register", h.RegisterUser)
		api.POST("/users/login", h.Login)
		api.POST("/users/logout", h.Logout)

		/* CITY */
		api.GET("/cities", h.GetCities)
	}
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./resources")
}

func (h *Handler) Error(ctx *gin.Context, code int, err error) {
	logrus.Error(err)

	ctx.JSON(code, gin.H{
		"error": err.Error(),
	})
}

func main() {
	router := gin.Default()

	router.GET("")
}

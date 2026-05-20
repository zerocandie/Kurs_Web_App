package handler

import (
	"net/http"
	"strconv"
	"time"

	"WebApp/internal/app/ds"

	"github.com/gin-gonic/gin"
)

// CreateEventInput модель запроса создания услуги
// @Description Данные для создания новой услуги/события
type CreateEventInput struct {
	Title       string `json:"title" example:"Rock Festival" binding:"required"`
	Description string `json:"description" example:"Annual rock music festival"`
	Image       string `json:"image" example:"festival_poster.jpg"`
	Video       string `json:"video" example:"promo.mp4"`
	EventDate   string `json:"event_date" example:"2026-12-31" binding:"required"`
	Location    string `json:"location" example:"Olympic Stadium"`
	CityID      uint   `json:"city_id" example:"1" binding:"required"`
}

// GetEvents godoc
// @Summary      Получить список услуг с фильтрацией
// @Tags         Events
// @Accept       json
// @Produce      json
// @Param        search     query  string  false  "Поиск по названию"
// @Param        city_id    query  uint    false  "ID города"
// @Param        date_from  query  string  false  "Дата от (YYYY-MM-DD)"
// @Param        date_to    query  string  false  "Дата до (YYYY-MM-DD)"
// @Success      200  {object}  map[string]interface{}  "Список услуг"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Router       /events [get]
func (h *Handler) GetEvents(ctx *gin.Context) {
	search := ctx.Query("search")
	cityID := ctx.Query("city_id")
	dateFrom := ctx.Query("date_from")
	dateTo := ctx.Query("date_to")

	events, err := h.Repository.GetEventsFiltered(search, cityID, dateFrom, dateTo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"count": len(events), "data": events})
}

// GetEventByID godoc
// @Summary      Получить услугу по ID
// @Tags         Events
// @Produce      json
// @Param        id  path  uint  true  "ID услуги"
// @Success      200  {object}  map[string]interface{}  "Данные услуги"
// @Failure      400  {object}  map[string]string  "Невалидный ID"
// @Failure      404  {object}  map[string]string  "Не найдено"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Router       /events/{id} [get]
func (h *Handler) GetEventByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}
	if idInt < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id must be positive"})
		return
	}

	event, err := h.Repository.GetEventByID(uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": event})
}

// CreateEvent godoc
// @Summary      Создать услугу
// @Tags         Events
// @Accept       json
// @Produce      json
// @Param        request  body  CreateEventInput  true  "Данные услуги"
// @Success      201  {object}  map[string]string  "Создано"
// @Failure      400  {object}  map[string]string  "Ошибка валидации"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Security     BearerAuth
// @Router       /events [post]
func (h *Handler) CreateEvent(ctx *gin.Context) {
	var input CreateEventInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	eventDate, err := time.Parse("2006-01-02", input.EventDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid event_date format"})
		return
	}

	event := ds.Event{
		Title: input.Title, Description: input.Description,
		Image: input.Image, Video: input.Video,
		EventDate: eventDate, Location: input.Location,
		CityID: input.CityID, IsActive: true,
	}
	if err := h.Repository.CreateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "data": event})
}
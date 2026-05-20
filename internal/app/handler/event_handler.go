package handler

import (
	"WebApp/internal/app/ds"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateEventInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Video       string `json:"video"`
	EventDate   string `json:"event_date"`
	Location    string `json:"location"`
	CityID      uint   `json:"city_id"`
}

/* ===================================
   GET ALL EVENTS + FILTER
   /api/events?search=xxx&city_id=1
=================================== */

func (h *Handler) GetEvents(ctx *gin.Context) {
	search := ctx.Query("search")
	cityID := ctx.Query("city_id")
	dateFrom := ctx.Query("date_from") // e.g., ?date_from=2026-05-01
	dateTo := ctx.Query("date_to")     // e.g., ?date_to=2026-12-31

	events, err := h.Repository.GetEventsFiltered(
		search,
		cityID,
		dateFrom,
		dateTo,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(events),
		"data":  events,
	})
}

/* ===================================
   GET EVENT BY ID
   /api/events/1
=================================== */

func (h *Handler) GetEventByID(
	ctx *gin.Context,
) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid id",
			},
		)
		return
	}

	// Explicitly convert int to uint
	if idInt < 0 {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid id: must be positive",
			},
		)
		return
	}
	id := uint(idInt)

	event, err := h.Repository.GetEventByID(id)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "event not found",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data": event,
		},
	)
}
/* ===================================
   CREATE EVENT
   POST /api/events
=================================== */

func (h *Handler) CreateEvent(
	ctx *gin.Context,
) {

	var input CreateEventInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if input.Title == "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "title is required",
			},
		)
		return
	}

	eventDate, err := time.Parse(
		"2006-01-02",
		input.EventDate,
	)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid event_date format, use YYYY-MM-DD",
			},
		)
		return
	}

	event := ds.Event{
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Video:       input.Video,
		EventDate:   eventDate,
		Location:    input.Location,
		CityID:      input.CityID,
		IsActive:    true,
	}

	err = h.Repository.CreateEvent(
		&event,
	)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"message": "event created",
			"data":    event,
		},
	)
}

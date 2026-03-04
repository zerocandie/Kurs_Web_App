package handler

import (
	"WebApp/internal/app/ds"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) ShowCreateEventForm(ctx *gin.Context) {

	cities, err := h.Repository.GetAllCities()
	if err != nil {
		logrus.Error("Failed to fetch cities:", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.HTML(http.StatusOK, "create_event.html", gin.H{
		"cities": cities,
	})
}
func (h *Handler) GetCitysChat(ctx *gin.Context) {
	var event []ds.Event
	var err error

	searchQuery := ctx.Query("query")

	if searchQuery == "" {
		event, err = h.Repository.GetAllChats()
	} else {
		event, err = h.Repository.GetEventsByTitle(searchQuery)
	}

	if err != nil {
		logrus.Error("Failed to fetch events:", err)
		ctx.String(http.StatusInternalServerError, "DB Error: %v", err)
		return
	}

	logrus.Infof("Fetched %d events", len(event))
	if len(event) > 0 {
		logrus.Infof("First event: %+v", event[0])
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"city":  event,
		"query": searchQuery,
	})
}

func (h *Handler) GetCityChat(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		logrus.Error("ID is empty")
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error("Invalid ID:", err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	event, err := h.Repository.GetAllChatsById(id)
	if err != nil {
		logrus.Error("Failed to fetch event:", err)
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.HTML(http.StatusOK, "order.html", gin.H{
		"event": event,
	})
}

func (h *Handler) AddEvent(ctx *gin.Context) {
	title := ctx.PostForm("title")
	about := ctx.PostForm("about")
	img := ctx.PostForm("img")
	location := ctx.PostForm("location")
	eventDateStr := ctx.PostForm("event_date")
	maxParticipantsStr := ctx.PostForm("max_participants")
	cityIDStr := ctx.PostForm("city_id")

	eventDate, err := time.Parse("2006-01-02", eventDateStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid date format")
		return
	}

	maxParticipants, err := strconv.Atoi(maxParticipantsStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid max participants")
		return
	}

	cityID, err := strconv.Atoi(cityIDStr)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid city ID")
		return
	}

	if title == "" {
		ctx.String(http.StatusBadRequest, "Title is required")
		return
	}

	event := ds.Event{
		Title:               title,
		About:               about,
		Img:                 img,
		Location:            location,
		EventDate:           eventDate,
		MaxParticipants:     maxParticipants,
		CurrentParticipants: 0,
		CityID:              uint(cityID),
		IsActive:            true,
	}
	createdEvent, err := h.Repository.CreateEvent(event)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to create event")
		return
	}
	ctx.Redirect(http.StatusFound, "/events/"+strconv.Itoa(int(createdEvent.ID)))

}

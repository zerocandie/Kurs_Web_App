package handler

import (
	"WebApp/internal/app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h Handler) GetCityItog(ctx *gin.Context) {
	Itogs, err := h.Repository.GetItogCity()
	if err != nil {
		logrus.Error(err)
	}
	ctx.HTML(http.StatusOK, "result.html", gin.H{
		"Itogs": Itogs,
	})
}

func (h *Handler) GetCitys(ctx *gin.Context) {
	var events []repository.Events_City
	var err error

	searchQuery := ctx.Query("query") // получаем значение из поля поиска
	if searchQuery == "" {            // если поле поиска пусто, то просто получаем из репозитория все записи
		events, err = h.Repository.GetEvents()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		events, err = h.Repository.GetCityByTitle(searchQuery) // в ином случае ищем заказ по заголовку
		if err != nil {
			logrus.Error(err)
		}
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"events": events,
		"query":  searchQuery, // передаем введенный запрос обратно на страницу
		// в ином случае оно будет очищаться при нажатии на кнопку
	})
}

func (h *Handler) GetCity(ctx *gin.Context) {

	idStr := ctx.Param("City")
	city, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error(err)
	}

	order, err := h.Repository.GetCity(city)
	if err != nil {
		logrus.Error(err)
	}
	ctx.HTML(http.StatusOK, "order.html", gin.H{
		"order": order,
	})
}

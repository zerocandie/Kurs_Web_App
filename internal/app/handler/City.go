package handler

import (
	"WebApp/internal/app/ds"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetCitys(ctx *gin.Context) {
	var citys []ds.City
	var err error

	searchQuery := ctx.Query("query")
	if searchQuery == "" {
		citys, err = h.Repository.GetCitys()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		citys, err = h.Repository.GetCitysByTitle(searchQuery)
		if err != nil {
			logrus.Error(err)
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"time":  time.Now().Format("15:04:05"),
			"citys": citys,
			"query": searchQuery,
		})
	}
}

func (h *Handler) GetCity(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error(err)
	}
	city, err := h.Repository.GetCity(id)
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "city.html", gin.H{
		"city": city,
	})

}

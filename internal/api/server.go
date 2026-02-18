package api

import (
	"WebApp/internal/app/handler"
	"WebApp/internal/app/repository"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	log.Println("Starting server...")
	repo, err := repository.NewRepository()
	if err != nil {
		logrus.Error("Ошибка инициализации репозитория")
	}

	handler := handler.NewHandler(repo)

	r := gin.Default()

	r.LoadHTMLGlob("Templates/*")
	r.Static("/static", "./resources")

	r.GET("/hello", handler.GetCitys)
	r.GET("/order/:id", handler.GetCity)
	r.GET("/result", handler.GetCityItog)
	r.Run()

	log.Println("Server down")
}

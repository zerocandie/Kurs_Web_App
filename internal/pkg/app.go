package pkg

import (
	"WebApp/internal/app/config"
	"WebApp/internal/app/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Config  *config.Config
	Router  *gin.Engine
	Handler *handler.Handler
}

func NewApp(conf *config.Config, router *gin.Engine, handler *handler.Handler) *Application {
	return &Application{
		Config:  conf,
		Router:  router,
		Handler: handler,
	}
}

// RunApp только запускает сервер. Регистрация роутов должна быть в main.go!
func (a *Application) RunApp() {
	logrus.Infof("🚀 Server starting on %s:%d", a.Config.ServiceHost, a.Config.ServicePort)

	err := a.Router.Run(fmt.Sprintf("%s:%d", a.Config.ServiceHost, a.Config.ServicePort))
	if err != nil {
		logrus.Fatal(err)
	}
}
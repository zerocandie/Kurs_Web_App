package main

// @title           E-Voting System API
// @version         1.0
// @description     Backend для системы электронного голосования с управлением услугами, заявками и пользователями
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     https://vk.com/bmstu_schedule
// @contact.email   support@evoting.ru

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api
// @schemes         http

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 JWT access token. Формат: `Bearer <token>`

import (
	// ✅ КРИТИЧЕСКИ ВАЖНО: импорт сгенерированной документации (с подчёркиванием!)
	_ "WebApp/cmd/WebApp/docs"

	"WebApp/internal/app/config"
	"WebApp/internal/app/dsn"
	"WebApp/internal/app/handler"
	"WebApp/internal/app/repository"
	"WebApp/internal/pkg"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	// Swagger middleware
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 1. Создаём роутер
	router := gin.Default()

	// ✅ 2. РЕГИСТРИРУЕМ SWAGGER *ПЕРВЫМ* (до любых других роутов!)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8082/swagger/doc.json"),
	))
	logrus.Info("📚 Swagger docs available at: http://localhost:8080/swagger/index.html")

	// 3. Инициализация (конфиг, БД, репозитории)
	conf, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("❌ config error: %v", err)
	}

	rep, err := repository.New(dsn.FromEnv())
	if err != nil {
		logrus.Fatalf("❌ repository error: %v", err)
	}

	hand := handler.NewHandler(rep)

	// ✅ 4. Регистрация ваших API-роутов (ПОСЛЕ Swagger, но ДО запуска)
	hand.RegisterRoutes(router)

	// 5. Создаём и запускаем приложение
	application := pkg.NewApp(conf, router, hand)
	application.RunApp()
}
package main

import (
	"WebApp/internal/app/ds"
	"WebApp/internal/app/dsn"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("[ERROR] Не удалось загрузить .env: %v\n", err)
		fmt.Println("[HINT] Убедись, что запускаешь команду ИЗ КОРНЯ проекта")
	} else {
		fmt.Println("[OK] .env loaded successfully")
	}

	// 2. Проверяем, что переменные действительно загрузились
	fmt.Printf("[DEBUG] DB_HOST='%s'\n", os.Getenv("DB_HOST"))
	fmt.Printf("[DEBUG] DB_PORT='%s'\n", os.Getenv("DB_PORT"))
	fmt.Printf("[DEBUG] DB_USER='%s'\n", os.Getenv("DB_USER"))
	fmt.Printf("[DEBUG] DB_PASS='%s'\n", os.Getenv("DB_PASS"))
	fmt.Printf("[DEBUG] DB_NAME='%s'\n", os.Getenv("DB_NAME"))

	dsnStr := dsn.FromEnv()
	fmt.Printf("[DEBUG] DSN: '%s'\n", dsnStr) // если пустой - увидишь ''

	// 4. Подключаемся
	db, err := gorm.Open(postgres.Open(dsnStr), &gorm.Config{})
	if err != nil {
		fmt.Printf("[ERROR] DB connect: %v\n", err)
		panic("failed to connect database")
	}
	// Migrate the schema
	err = db.AutoMigrate(
		&ds.Event{},
	)
	if err != nil {
		panic("cant migrate db")
	}

	fmt.Println("migrate success")
}

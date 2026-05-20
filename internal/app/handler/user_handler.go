package handler

import (
	"WebApp/internal/app/auth"
	"WebApp/internal/app/ds"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterInput модель регистрации
// @Description Данные для создания аккаунта
type RegisterInput struct {
	Login    string `json:"login" example:"john_doe" binding:"required,min=3,max=100"`
	Email    string `json:"email" example:"user@example.com" binding:"required,email"`
	Password string `json:"password" example:"SecurePass123!" binding:"required,min=8"`
}

// LoginInput модель аутентификации
// @Description Учетные данные для входа
type LoginInput struct {
	Email    string `json:"email" example:"user@example.com" binding:"required,email"`
	Password string `json:"password" example:"SecurePass123!" binding:"required"`
}

// RegisterUser godoc
// @Summary      Регистрация
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body  RegisterInput  true  "Данные"
// @Success      201  {object}  map[string]string  "Пользователь создан"
// @Failure      400  {object}  map[string]string  "Ошибка валидации"
// @Failure      500  {object}  map[string]string  "Ошибка сервера"
// @Router       /users/register [post]
func (h *Handler) RegisterUser(ctx *gin.Context) {
	var input RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "password hash error"})
		return
	}
	user := ds.User{Login: input.Login, Email: input.Email, Password: string(hashedPassword), Role: ds.RoleCitizen}
	if err := h.Repository.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

// Login godoc
// @Summary      Аутентификация
// @Description  Выдаёт JWT и устанавливает httpOnly cookie. Токен также возвращается в теле ответа для использования в Postman/Insomnia.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body  LoginInput  true  "Email и пароль"
// @Success      200  {object}  LoginResponse  "Токен и данные пользователя"
// @Failure      400  {object}  map[string]string  "Ошибка валидации"
// @Failure      401  {object}  map[string]string  "Неверные учетные данные"
// @Router       /auth/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	var input LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Repository.GetUserByEmail(input.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	ctx.SetCookie("token", token, 3600*24, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"login": user.Login,
			"role":  user.Role,
		},
	})
}

// Logout godoc
// @Summary      Выход
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  map[string]string  "Выход успешен"
// @Security     BearerAuth
// @Router       /auth/logout [post]
func (h *Handler) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
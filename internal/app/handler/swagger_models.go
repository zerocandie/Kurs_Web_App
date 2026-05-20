// internal/app/handler/swagger_models.go
package handler

// SuccessMessage простой ответ с сообщением
// @Description Используется для операций без возвращаемых данных
type SuccessMessage struct {
	Message string `json:"message" example:"operation successful"`
}

// ErrorResponse стандартный формат ошибки API
// @Description Возвращается при любой ошибке сервера или валидации
type ErrorResponse struct {
	Error string `json:"error" example:"invalid request"`
}

// CartResponse модель иконки корзины
// @Description Информация о черновике заявки текущего пользователя
type CartResponse struct {
	OrderID *uint `json:"order_id" example:"5"`
	Count   int   `json:"count" example:"3"`
}

// VoteResponse модель ответа заявки с услугами
// @Description Детальная информация о заявке включая связанные услуги
type VoteResponse struct {
	Data interface{} `json:"data"` // замените на ds.Vote при необходимости
}

// VotesListResponse модель ответа списка заявок
// @Description Список заявок с фильтрацией и вычисляемыми полями
type VotesListResponse struct {
	Count int           `json:"count" example:"10"`
	Data  []interface{} `json:"data"` // замените на []ds.Vote при необходимости
}

// EventsListResponse модель ответа списка услуг
// @Description Пагинированный список событий с метаданными
type EventsListResponse struct {
	Count int           `json:"count" example:"15"`
	Data  []interface{} `json:"data"` // замените на []ds.Event при необходимости
}

// LoginResponse модель успешного входа
// @Description Ответ с токеном и данными пользователя
type LoginResponse struct {
	Message string `json:"message" example:"login success"`
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User    struct {
		ID    uint   `json:"id" example:"1"`
		Login string `json:"login" example:"john_doe"`
		Email string `json:"email" example:"user@example.com"`
		Role  string `json:"role" example:"citizen"`
	} `json:"user"`
}

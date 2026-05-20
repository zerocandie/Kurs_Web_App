package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// contextKey тип для ключей контекста
type contextKey string

const (
	// SessionKey ключ для хранения сессии в контексте
	SessionKey contextKey = "session"
)

// Middleware структура middleware
type Middleware struct {
	sessionManager *SessionManager
}

// NewMiddleware создает новый middleware
func NewMiddleware(sessionManager *SessionManager) *Middleware {
	return &Middleware{
		sessionManager: sessionManager,
	}
}

// Auth middleware для проверки аутентификации
func (m *Middleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка Authorization или Cookie
		token := m.getTokenFromRequest(r)

		if token == "" {
			http.Error(w, `{"error": "unauthorized"}`, http.StatusUnauthorized)
			return
		}

		// Проверяем сессию
		session, err := m.sessionManager.GetSession(token)
		if err != nil {
			http.Error(w, `{"error": "invalid or expired session"}`, http.StatusUnauthorized)
			return
		}

		// Добавляем сессию в контекст
		ctx := context.WithValue(r.Context(), SessionKey, session)
		next(w, r.WithContext(ctx))
	}
}

// getTokenFromRequest извлекает токен из запроса
func (m *Middleware) getTokenFromRequest(r *http.Request) string {
	// Пробуем получить из заголовка Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}

	// Пробуем получить из Cookie
	cookie, err := r.Cookie("session_id")
	if err == nil {
		return cookie.Value
	}

	// Пробуем получить из query параметра
	token := r.URL.Query().Get("token")
	if token != "" {
		return token
	}

	return ""
}

// GetSessionFromContext получает сессию из контекста
func GetSessionFromContext(ctx context.Context) (*Session, error) {
	session, ok := ctx.Value(SessionKey).(*Session)
	if !ok {
		return nil, fmt.Errorf("session not found in context")
	}
	return session, nil
}

// CORS middleware для обработки CORS
func (m *Middleware) CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// Logger middleware для логгирования запросов
func (m *Middleware) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логгируем запрос (в реальном приложении используйте logger)
		logrus.Infof("Request: %s %s", r.Method, r.URL.Path);

		// Вызываем следующий handler
		next(w, r)
	}
}

// RateLimit простой rate limiting middleware
func (m *Middleware) RateLimit(maxRequests int, window time.Duration) func(http.HandlerFunc) http.HandlerFunc {
	requests := make(map[string][]time.Time)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			now := time.Now()

			// Очищаем старые запросы
			if reqs, exists := requests[ip]; exists {
				var validReqs []time.Time
				for _, t := range reqs {
					if now.Sub(t) < window {
						validReqs = append(validReqs, t)
					}
				}
				requests[ip] = validReqs
			}

			// Проверяем лимит
			if len(requests[ip]) >= maxRequests {
				http.Error(w, `{"error": "rate limit exceeded"}`, http.StatusTooManyRequests)
				return
			}

			// Добавляем запрос
			requests[ip] = append(requests[ip], now)

			next(w, r)
		}
	}
}
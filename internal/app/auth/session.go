package auth

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Session структура сессии
type Session struct {
	ID        string                 `json:"id"`
	UserID    int                    `json:"user_id"`
	Username  string                 `json:"username"`
	CreatedAt time.Time              `json:"created_at"`
	ExpiresAt time.Time              `json:"expires_at"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// SessionManager менеджер сессий
type SessionManager struct {
	redis   *RedisClient
	timeout time.Duration
}

// NewSessionManager создает новый менеджер сессий
func NewSessionManager(redis *RedisClient, timeout time.Duration) *SessionManager {
	return &SessionManager{
		redis:   redis,
		timeout: timeout,
	}
}

// CreateSession создает новую сессию
func (sm *SessionManager) CreateSession(userID int, username string) (*Session, error) {
	sessionID := uuid.New().String()
	now := time.Now()

	session := &Session{
		ID:        sessionID,
		UserID:    userID,
		Username:  username,
		CreatedAt: now,
		ExpiresAt: now.Add(sm.timeout),
		Data:      make(map[string]interface{}),
	}

	// Сериализуем сессию в JSON
	sessionData, err := json.Marshal(session)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal session: %w", err)
	}

	// Сохраняем в Redis
	key := fmt.Sprintf("session:%s", sessionID)
	if err := sm.redis.Set(key, sessionData, sm.timeout); err != nil {
		return nil, fmt.Errorf("failed to save session: %w", err)
	}

	return session, nil
}

// GetSession получает сессию по ID
func (sm *SessionManager) GetSession(sessionID string) (*Session, error) {
	key := fmt.Sprintf("session:%s", sessionID)

	data, err := sm.redis.Get(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	var session Session
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	// Проверяем не истекла ли сессия
	if time.Now().After(session.ExpiresAt) {
		sm.DeleteSession(sessionID)
		return nil, fmt.Errorf("session expired")
	}

	// Продлеваем сессию
	session.ExpiresAt = time.Now().Add(sm.timeout)
	updatedData, _ := json.Marshal(session)
	sm.redis.Set(key, updatedData, sm.timeout)

	return &session, nil
}

// DeleteSession удаляет сессию
func (sm *SessionManager) DeleteSession(sessionID string) error {
	key := fmt.Sprintf("session:%s", sessionID)
	return sm.redis.Delete(key)
}

// SetSessionData устанавливает произвольные данные в сессию
func (sm *SessionManager) SetSessionData(sessionID, key string, value interface{}) error {
	session, err := sm.GetSession(sessionID)
	if err != nil {
		return err
	}

	session.Data[key] = value

	sessionData, err := json.Marshal(session)
	if err != nil {
		return err
	}

	redisKey := fmt.Sprintf("session:%s", sessionID)
	return sm.redis.Set(redisKey, sessionData, sm.timeout)
}

// GetSessionData получает данные из сессии
func (sm *SessionManager) GetSessionData(sessionID, key string) (interface{}, error) {
	session, err := sm.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	value, exists := session.Data[key]
	if !exists {
		return nil, fmt.Errorf("key %s not found in session", key)
	}

	return value, nil
}

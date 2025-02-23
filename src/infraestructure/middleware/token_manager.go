package middleware

import (
	"sync"
	"time"
)

type TokenManager struct {
	userSecrets map[string]string
	mutex       sync.RWMutex
}

var tokenManager = &TokenManager{userSecrets: make(map[string]string)}

func generateUserSecret(userID string) string {
	return userID + "_" + time.Now().Format("20060102150405")
}

func GetUserSecret(userID string) string {
	tokenManager.mutex.RLock()
	secret, exists := tokenManager.userSecrets[userID]
	tokenManager.mutex.RUnlock()

	if exists {
		return secret
	}

	newSecret := generateUserSecret(userID)
	SetUserSecret(userID, newSecret)
	return newSecret
}

func SetUserSecret(userID string, secret string) {
	tokenManager.mutex.Lock()
	tokenManager.userSecrets[userID] = secret
	tokenManager.mutex.Unlock()
}

func InvalidateUserTokens(userID string) {
	tokenManager.mutex.Lock()
	tokenManager.userSecrets[userID] = generateUserSecret(userID)
	tokenManager.mutex.Unlock()
}

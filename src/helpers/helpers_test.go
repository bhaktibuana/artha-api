package helpers_test

import (
	"artha-api/src/configs"
	"artha-api/src/helpers"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	// Mock JWT payload
	payload := jwt.MapClaims{
		"sub": "user123",
		"exp": time.Now().Add(1 * time.Hour).Unix(), // Set expiry time to 1 hour from now
	}

	// Generate JWT token
	token, err := helpers.GenerateJWT(payload, 1*time.Hour)
	assert.NoError(t, err, "GenerateJWT should not return an error")

	// Parse and verify the JWT token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.AppConfig().JWT_SECRET_KEY), nil
	})
	assert.NoError(t, err, "jwt.Parse should not return an error")

	// Verify token claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok, "Token claims should be of type jwt.MapClaims")
	assert.True(t, parsedToken.Valid, "JWT token should be valid")
	assert.Equal(t, payload["sub"], claims["sub"], "JWT token subject should match")
}

func TestHashPassword(t *testing.T) {
	// Test case: Hashing succeeds
	t.Run("HashingSucceeds", func(t *testing.T) {
		password := "secretpassword"
		hashedPassword := helpers.HashPassword(password)

		// You might want to adjust this expectation based on the actual hashed value
		expectedHash := "0d3418706dcc4c2bf745542eb58a34f2bb8e4ae6bdd24d195dda646108551474"

		assert.Equal(t, expectedHash, hashedPassword, "Hashed password should match expected value")
	})
}

func TestGenerateAccountId(t *testing.T) {
	randomUsername, randomTag, randomAccountId := helpers.GenerateAccountId()

	// Check if randomUsername is not empty
	if randomUsername == "" {
		t.Errorf("Random username is empty")
	}

	// Check if randomTag is not empty
	if randomTag == "" {
		t.Errorf("Random tag is empty")
	}

	// Check if randomAccountId is in the correct format
	parts := strings.Split(randomAccountId, "#")
	if len(parts) != 2 {
		t.Errorf("Random account ID is not in the correct format: %s", randomAccountId)
	}
}

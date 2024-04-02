package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateAccountId Helper
/*
 * @returns (string, string, string) (username, tag, accountID)
 */
func GenerateAccountId() (string, string, string) {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	usernameLength := rand.Intn(16) + 1
	randomUsername := make([]byte, usernameLength)
	for i := 0; i < usernameLength; i++ {
		randomUsername[i] = charset[rand.Intn(len(charset))]
	}
	randomUsernameString := string(randomUsername)

	tagLineLength := rand.Intn(5) + 1
	randomTag := make([]byte, tagLineLength)
	for i := 0; i < tagLineLength; i++ {
		randomTag[i] = charset[rand.Intn(len(charset))]
	}
	randomTagString := string(randomTag)

	randomAccountId := fmt.Sprintf("%s#%s", randomUsernameString, randomTagString)

	return randomUsernameString, randomTagString, randomAccountId
}

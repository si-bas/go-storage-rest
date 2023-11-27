package shared

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"
)

func GetContextValueAsString(ctx context.Context, key string) string {
	val, ok := ctx.Value(key).(string)
	if ok {
		return val
	}

	return ""
}

func GetContextValueAsNumber(ctx context.Context, key string) (uint64, error) {
	str := GetContextValueAsString(ctx, key)

	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func GenerateSlug(input string) string {
	// Convert the string to lowercase
	slug := strings.ToLower(input)

	// Remove special characters and replace spaces with dashes
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")

	// Remove leading and trailing dashes
	slug = strings.Trim(slug, "-")

	return slug
}

func GenerateAPIKey() (string, error) {
	length := 30

	// The length parameter specifies the number of bytes for the key
	keyBytes := make([]byte, length)
	_, err := rand.Read(keyBytes)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a hex-encoded string
	apiKey := hex.EncodeToString(keyBytes)

	return apiKey, nil
}

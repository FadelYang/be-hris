package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func StringToUUID(s string) (uuid.UUID, error) {
	formattedUUID, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, err
	}

	return formattedUUID, nil
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func GenerateSlug(input string) string {
	slug := strings.ToLower(input)

	// Replace spaces with -
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove non alphanumeric except -
	re := regexp.MustCompile(`[^a-z0-9-]+`)
	slug = re.ReplaceAllString(slug, "")

	// Remove duplicate -
	re = regexp.MustCompile(`-+`)
	slug = re.ReplaceAllString(slug, "-")

	return strings.Trim(slug, "-")
}

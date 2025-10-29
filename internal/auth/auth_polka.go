package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := strings.TrimSpace(headers.Get("Authorization"))
	if apiKey == "" {
		return "", errors.New("missing api key")
	}
	scheme, rest, ok := strings.Cut(apiKey, " ")

	if !ok || !strings.EqualFold(scheme, "ApiKey") {
		return "", errors.New("authorization scheme must be Bearer")
	}

	key := strings.TrimSpace(rest)

	if key == "" {
		return "", errors.New("empty api key")
	}
	return key, nil
}


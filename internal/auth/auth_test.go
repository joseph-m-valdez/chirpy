package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func makeTestToken(t *testing.T, sub string, secret string, expOffset time.Duration) string {
	t.Helper()
	claims := &jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expOffset)),
		Subject:   sub,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("sign token: %v", err)
	}
	return s
}

func TestValidateJWT_RoundTripOK(t *testing.T) {
	secret := "supersecret"
	id := uuid.New()
	// not expired
	token := makeTestToken(t, id.String(), secret, 5*time.Minute)

	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT returned error: %v", err)
	}
	if gotID != id {
		t.Fatalf("want userID %s, got %s", id, gotID)
	}
}


func TestValidateJWT_Expired(t *testing.T) {
	secret := "supersecret"
	id := uuid.New()
	// already expired (in the past)
	token := makeTestToken(t, id.String(), secret, -1*time.Minute)

	_, err := ValidateJWT(token, secret)
	if err == nil {
		t.Fatalf("expected error for expired token, got nil")
	}
}

func TestValidateJWT_WrongSecret(t *testing.T) {
	id := uuid.New()
	token := makeTestToken(t, id.String(), "right-secret", 5*time.Minute)

	_, err := ValidateJWT(token, "wrong-secret")
	if err == nil {
		t.Fatalf("expected signature verification error, got nil")
	}
}

func TestValidateJWT_InvalidSubject(t *testing.T) {
	secret := "supersecret"
	// sub is not a UUID
	token := makeTestToken(t, "not-a-uuid", secret, 5*time.Minute)

	_, err := ValidateJWT(token, secret)
	if err == nil {
		t.Fatalf("expected error for invalid subject, got nil")
	}
}

func TestMakeJWTAndValidate(t *testing.T) {
	secret := "supersecret"
	id := uuid.New()

	token, err := MakeJWT(id, secret, 2*time.Minute)
	if err != nil {
		t.Fatalf("MakeJWT error: %v", err)
	}

	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT error: %v", err)
	}
	if gotID != id {
		t.Fatalf("want %s, got %s", id, gotID)
	}
}



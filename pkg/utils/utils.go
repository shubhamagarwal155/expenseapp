package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserId int
	jwt.RegisteredClaims
}

var secretKey = "sdvhvsadhvshavdhjsavjdhvsahvdasvdyvywwwwww131231208312iug3iu12gi123viy123vuiv312uyv3123yvyuipv"

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, statuscode int, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)

	return json.NewEncoder(w).Encode(payload)
}

func GenerateJWTAccessToken(userId int) (string, error) {
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func ValidateJWTToken(jwtToken string) (claims *Claims, msg string) {
	token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(t *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })

	if err == nil {
		claims, ok := token.Claims.(*Claims)

		if !ok {
			return nil, "invalid token"
		} else {
			return claims, ""
		}
	} else {
		return nil, "error parsing token"
	}
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

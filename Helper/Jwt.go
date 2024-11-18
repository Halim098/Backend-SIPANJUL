package Helper

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey = func() []byte {
    key := os.Getenv("JWT_PRIVATE_KEY")
    if key == "" {
        panic("JWT_PRIVATE_KEY is not set in environment variables")
    }
    return []byte(key)
}()

func GenerateJWT(id uint) (string, error) {
    tokenTTL, err := strconv.Atoi(os.Getenv("TOKEN_TTL"))
    if err != nil || tokenTTL <= 0 {
        return "", fmt.Errorf("invalid TOKEN_TTL value")
    }

    claims := jwt.MapClaims{
        "id":  id,
        "exp": time.Now().Add(time.Minute * time.Duration(tokenTTL)).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(privateKey)
}
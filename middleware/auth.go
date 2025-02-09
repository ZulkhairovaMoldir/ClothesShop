package middleware

import (
    "fmt"
    "log"
    "github.com/golang-jwt/jwt/v4"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "strings"
    "time"
)

var secretKey []byte

func init() {
    secretKey = getSecretKey()
}

func getSecretKey() []byte {
    key := os.Getenv("SECRET_KEY")
    if key == "" {
        log.Fatal("SECRET_KEY environment variable not set")
    }
    return []byte(key)
}

func GenerateJWT(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "userID": userID,
        "exp":    time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secretKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
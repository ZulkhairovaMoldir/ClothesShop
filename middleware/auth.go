package middleware

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "github.com/golang-jwt/jwt/v4"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

var secretKey = generateSecretKey()

func generateSecretKey() []byte {
    key := make([]byte, 32) // 32 bytes = 256 bits
    _, err := rand.Read(key)
    if err != nil {
        panic(err)
    }
    return key
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
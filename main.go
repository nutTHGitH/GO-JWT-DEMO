package main

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        payload, err := ValidateToken(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT: " + err.Error()})
            return
        }

        ctx := StorePayload(c.Request.Context(), payload)
        c.Request = c.Request.WithContext(ctx)
        c.Next()
    }
}

func main() {
    router := gin.Default()

    router.GET("/api/test", JwtMiddleware(), func(c *gin.Context) {
        payload := GetPayload(c.Request.Context())
        if payload == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No JWT payload found"})
            return
        }
        c.JSON(http.StatusOK, payload)
    })

    fmt.Println("Server started at http://localhost:8080")
    router.Run(":8080")
}

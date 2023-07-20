package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joshuaetim/gallery_go/handler"
)

func AuthorizeJWT(strict ...bool) gin.HandlerFunc {
	if len(strict) == 0 {
		strict = append(strict, true)
	}
	return func(ctx *gin.Context) {
		var BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			if strict[0] {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "no auth header found",
				})
			}
			return
		}
		// take care of empty header contigency
		if len(BearerSchema) >= len(authHeader) {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "invalid header entry",
			})
			return
		}
		tokenString := authHeader[len(BearerSchema):]

		token, err := handler.ValidateToken(tokenString)
		if err != nil {
			if strict[0] {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid token",
				})
			}
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "wrong claims type"})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		ctx.Set("userID", claims["userID"])
		fmt.Println("during auth: ", claims["userID"])
	}
}

package middleware

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var jwtSecret = []byte("your-secret-key")

func JWTMiddleware(ctx *context.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, iris.NewProblem().Status(iris.StatusUnauthorized).Title("Invalid token")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		return
	}

	ctx.Values().Set("user_id", claims["user_id"])
	ctx.Values().Set("role", claims["role"])

	ctx.Next()
}

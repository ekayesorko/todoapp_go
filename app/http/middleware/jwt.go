package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"todolist-assignment/app/serializer"
	"todolist-assignment/utils"
)

func JWTParser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token, err := utils.GetTokenFromContext(ctx)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, err)
		}
		userID, err2 := utils.VerifyToken(token)
		if err2 != nil {
			return ctx.JSON(http.StatusUnauthorized, err)
		}
		ctx.Set("LoggedInUser", &serializer.LoggedInUser{
			UserID: userID,
		})
		return next(ctx)
	}
}

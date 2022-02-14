package utils

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"strings"
	"todolist-assignment/infra/errors"
)

type claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func StructToStruct(input interface{}, output interface{}) error {
	if b, err := json.Marshal(input); err == nil {
		return json.Unmarshal(b, &output)
	} else {
		return err
	}
}

func GetTokenFromContext(ctx echo.Context) (string, *errors.RestErr) {
	authHead, err := ctx.Request().Header["Authorization"]
	if err == false {
		return "", errors.NewUnauthorizedError("no jwt")
	}
	token := strings.Split(authHead[0], " ")[1]
	return token, nil
}

func VerifyToken(bearToken string) (uint, *errors.RestErr) {
	claims := &claims{}
	tkn, err := jwt.ParseWithClaims(bearToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte("secret-key"), nil
		},
	)
	if err != nil {
		return 0, errors.NewUnauthorizedError("json parse error")
	}
	if !tkn.Valid {
		return 0, errors.NewUnauthorizedError("invalid jwt signature")
	}
	return claims.UserID, nil
}

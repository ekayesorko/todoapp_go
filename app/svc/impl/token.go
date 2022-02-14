package impl

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"todolist-assignment/app/repository"
	"todolist-assignment/app/svc"
	"todolist-assignment/infra/errors"
)

type token struct {
	uRep repository.IUsers
}

type claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func (t *token) CreateToken(_userID uint) (string, *errors.RestErr) {
	expirationTime := time.Now().Add(50 * time.Minute).Unix()
	atMap := &claims{
		UserID: _userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atMap)
	token, err := at.SignedString([]byte("secret-key"))
	if err != nil {
		return "", &errors.RestErr{Message: err.Error()}
	}
	return token, nil
}

func (t *token) VerifyToken(bearToken string) (uint, *errors.RestErr) {
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

func NewTokenService(_ur repository.IUsers) svc.IToken {
	return &token{
		uRep: _ur,
	}
}

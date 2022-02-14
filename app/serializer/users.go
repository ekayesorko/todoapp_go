package serializer

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignupResp struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type SignupReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (sr SignupReq) Validate() error {
	return validation.ValidateStruct(&sr,
		validation.Field(&sr.Email, validation.Required, is.EmailFormat),
		validation.Field(&sr.Password, validation.Required),
		validation.Field(&sr.Username, validation.Required, validation.Length(5, 10)),
	)
}

package auth

import (
	"fmt"

	"github.com/Allan-Nava/gocloak-ddd/utils"
	"github.com/go-playground/validator"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

//
var validate = validator.New()

func ValidateStruct(request LoginRequest) *utils.ApiError {
	var errorReturn *utils.ApiError
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element utils.ApiError
			element.Message = fmt.Sprint(err) //err.StructNamespace()
			element.Response = "KO"
			errorReturn = &element
		}
	}
	return errorReturn
}

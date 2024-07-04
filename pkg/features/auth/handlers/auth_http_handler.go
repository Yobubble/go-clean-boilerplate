package handlers

import (
	"fmt"
	"net/http"

	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/constants"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/models"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/usecases"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils/commons"
	"github.com/gin-gonic/gin"
)

type authHttpHandler struct {
	usecase usecases.AuthUseCase
}

// SignUp implements AuthHandler.
func (a *authHttpHandler) SignUp(c *gin.Context) {
	fmt.Println(c.FullPath(), "SignUp")

	var signUpBody models.SignUpBodyModel

	if err := c.BindJSON(&signUpBody); err != nil {
		commons.Response(c, nil, commons.BodyParseError, http.StatusBadRequest, err.Error())
	}

	if err := a.usecase.InsertNewUser(signUpBody); err != nil {
		commons.Response(c, nil, constants.InsertUserError, http.StatusInternalServerError, err.Error())
	}

	commons.Response(c, nil, constants.InsertNewUserSuccess, http.StatusOK, "")
}

// SignIn implements Handler.
func (a *authHttpHandler) SignIn(c *gin.Context) {
	fmt.Println(c.FullPath(), "SignIn")

	var signInBody models.SignInBodyModel

	if err := c.BindJSON(&signInBody); err != nil {
		commons.Response(c, nil, commons.BodyParseError, http.StatusBadRequest, err.Error())
	}

	if token, err := a.usecase.ValidateUser(signInBody); err != nil {
		commons.Response(c, nil, constants.ValidateUserError, http.StatusUnauthorized, err.Error())
	} else {
		commons.Response(c, token, "", http.StatusOK, "")
	}
}

func NewAuthHttpHandler(usecase usecases.AuthUseCase) AuthHandler {
	return &authHttpHandler{
		usecase: usecase,
	}
}
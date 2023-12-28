package http

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/labstack/echo"

	"github.com/alifahsanilsatria/twitter-clone/domain"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	userUsecase domain.UserUsecase
	logger      *logrus.Logger
}

func NewUserHandler(
	e *echo.Echo,
	us domain.UserUsecase,
	logger *logrus.Logger,
) {
	handler := &UserHandler{
		userUsecase: us,
		logger:      logger,
	}
	e.POST("/sign-up", handler.SignUp)
}

// Store will store the article by given request body
func (handler *UserHandler) SignUp(c echo.Context) error {
	logData := logrus.Fields{
		"method": "handler.SignUp",
	}
	var reqPayload domain.SignUpParam
	errParsingReqPayload := c.Bind(&reqPayload)
	if errParsingReqPayload != nil {
		logData["error_parsing_request_payload"] = errParsingReqPayload.Error()
		handler.logger.
			WithFields(logData).
			WithError(errParsingReqPayload).
			Errorln("error when parsing request payload")
		return c.JSON(http.StatusUnprocessableEntity, errParsingReqPayload.Error())
	}

	logData["request_payload"] = fmt.Sprintf("%+v", reqPayload)

	errvalidateSignUpParam := validateSignUpParam(reqPayload)
	if errParsingReqPayload != nil {
		logData["error_validate_sign_up_param"] = errvalidateSignUpParam.Error()
		handler.logger.
			WithFields(logData).
			WithError(errvalidateSignUpParam).
			Errorln("error when validate sign up param")
		return c.JSON(http.StatusBadRequest, ResponseError{Message: errvalidateSignUpParam.Error()})
	}

	ctx := c.Request().Context()
	signUpUsecaseResult, errorSignUpUsecase := handler.userUsecase.SignUp(ctx, reqPayload)
	if errorSignUpUsecase != nil {
		logData["error_sign_up_usecase"] = errorSignUpUsecase.Error()
		handler.logger.
			WithFields(logData).
			WithError(errorSignUpUsecase).
			Errorln("error when parsing request payload")
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: errorSignUpUsecase.Error()})
	}

	return c.JSON(http.StatusOK, signUpUsecaseResult)
}

func validateSignUpParam(param domain.SignUpParam) error {
	if param.Username == "" || param.Password == "" || param.Email == "" || param.CompleteName == "" {
		if param.Username == "" {
			return errors.New("username is empty")
		} else if param.Password == "" {
			return errors.New("password is empty")
		} else if param.Email == "" {
			return errors.New("email is empty")
		} else {
			return errors.New("complete_name is empty")
		}
	}

	usernameRegex, _ := regexp.Compile("^[A-Za-z0-9_]{5,15}$")
	isUsernameMatch := usernameRegex.Match([]byte(param.Username))
	if !isUsernameMatch {
		return errors.New("your username must be more than 4 characters long, can be up to 15 characters, and can only contain letters, numbers, and underscores. no spaces are allowed")
	}

	regexContainUpperCaseLetters, _ := regexp.Compile("[A-Z]")
	regexContainLowerCaseLetters, _ := regexp.Compile("[a-z]")
	regexContainNumbers, _ := regexp.Compile("[0-9]")
	regexContainSymbols, _ := regexp.Compile("[,.\\/!@#$%^&*()_+-=';:]")

	isPasswordMatch := len(param.Password) >= 12 && regexContainUpperCaseLetters.Match([]byte(param.Password)) && regexContainLowerCaseLetters.Match([]byte(param.Password)) && regexContainNumbers.Match([]byte(param.Password)) && regexContainSymbols.Match([]byte(param.Password))
	if !isPasswordMatch {
		return errors.New("your password must be at least 12 characters long and use a mix of uppercase, lowercase, numbers, and symbols character")
	}

	emailRegex, _ := regexp.Compile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
	isEmailMatch := emailRegex.Match([]byte(param.Email))
	if !isEmailMatch {
		return errors.New("your email is in invalid form")
	}

	if len(param.CompleteName) > 50 {
		return errors.New("your display name must be maximum 50 characters long")
	}

	return nil
}

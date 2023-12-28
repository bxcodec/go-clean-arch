package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/alifahsanilsatria/twitter-clone/common"
	"github.com/alifahsanilsatria/twitter-clone/domain"
	"github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepository domain.UserRepository
	logger         *logrus.Logger
}

func NewUserUsecase(
	userRepository domain.UserRepository,
	logger *logrus.Logger,
) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (uc *userUsecase) SignUp(ctx context.Context, param domain.SignUpParam) (domain.SignUpResult, error) {
	logData := logrus.Fields{
		"method": "userUsecase.SignUp",
		"param":  fmt.Sprintf("%+v", param),
	}
	getUserByUsernameOrEmailParam := domain.GetUserByUsernameOrEmailParam{
		Username: strings.ToLower(param.Username),
		Email:    param.Email,
	}
	getUserByUsernameOrEmailResp, errGetUserByUsernameOrEmail := uc.userRepository.GetUserByUsernameOrEmail(ctx, getUserByUsernameOrEmailParam)
	if errGetUserByUsernameOrEmail != nil {
		logData["error_get_user_by_username_or_email"] = errGetUserByUsernameOrEmail.Error()
		uc.logger.
			WithFields(logData).
			WithError(errGetUserByUsernameOrEmail).
			Errorln("error on GetUserByUsernameOrEmail")
		return domain.SignUpResult{}, errGetUserByUsernameOrEmail
	}

	if getUserByUsernameOrEmailResp.Id > 0 {
		err := errors.New("username or email already exists")
		return domain.SignUpResult{}, err
	}

	bytesHashedPassword, errHashPassword := common.HashPassword(param.Password)
	if errHashPassword != nil {
		logData["error_hash_password"] = errHashPassword.Error()
		uc.logger.
			WithFields(logData).
			WithError(errHashPassword).
			Errorln("error on HashPassword")
		return domain.SignUpResult{}, errHashPassword
	}

	hashedPassword := string(bytesHashedPassword)

	createNewUserAccountParam := domain.CreateNewUserAccountParam{
		Username:       param.Username,
		HashedPassword: hashedPassword,
		Email:          param.Email,
		CompleteName:   param.CompleteName,
	}

	createNewUserAccountResp, errCreateNewUserAccount := uc.userRepository.CreateNewUserAccount(ctx, createNewUserAccountParam)
	if errCreateNewUserAccount != nil {
		logData["error_create_new_user_account"] = errCreateNewUserAccount.Error()
		uc.logger.
			WithFields(logData).
			WithError(errCreateNewUserAccount).
			Errorln("error on CreateNewUserAccount")
		return domain.SignUpResult{}, errCreateNewUserAccount
	}

	response := domain.SignUpResult{
		Id: createNewUserAccountResp.Id,
	}

	return response, nil

}

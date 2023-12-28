package db

import (
	"context"
	"fmt"
	"time"

	commonWrapper "github.com/alifahsanilsatria/twitter-clone/common/wrapper"
	"github.com/alifahsanilsatria/twitter-clone/domain"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db     commonWrapper.SQLWrapper
	logger *logrus.Logger
}

func NewUserRepository(
	db commonWrapper.SQLWrapper,
	logger *logrus.Logger,
) domain.UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (repo *userRepository) GetUserByUsernameOrEmail(ctx context.Context, param domain.GetUserByUsernameOrEmailParam) (domain.GetUserByUsernameOrEmailResult, error) {
	logData := logrus.Fields{
		"method": "userRepository.GetUserByUsernameOrEmail",
		"param":  fmt.Sprintf("%+v", param),
	}
	query := `
		select id
		from user u
		where username = $1
		or email = $2
	`

	args := []interface{}{
		param.Username,
		param.Email,
	}

	queryRowContextResp := repo.db.QueryRowContext(ctx, query, args...)

	response := domain.GetUserByUsernameOrEmailResult{}
	errScan := queryRowContextResp.Scan(&response.Id)
	if errScan != nil {
		logData["error_scan"] = errScan.Error()
		repo.logger.
			WithFields(logData).
			WithError(errScan).
			Errorln("error on scan")
	} else {
		logData["response"] = fmt.Sprintf("%+v", response)
		repo.logger.
			WithFields(logData).
			WithError(errScan).
			Debugln("success get response")
	}

	return response, errScan
}

func (repo *userRepository) CreateNewUserAccount(ctx context.Context, param domain.CreateNewUserAccountParam) (domain.CreateNewUserAccountResult, error) {
	logData := logrus.Fields{
		"method": "userRepository.CreateNewUserAccount",
		"param":  fmt.Sprintf("%+v", param),
	}

	query := `
		insert into user
		(username, password, email, complete_name, is_deleted, created_at)
		values
		($1, $2, $3, $4, $5, $6, $7)
	`

	args := []interface{}{
		param.Username,
		param.HashedPassword,
		param.Email,
		param.CompleteName,
		false,
		time.Now(),
	}

	execContextResp, errQuery := repo.db.ExecContext(ctx, query, args...)
	if errQuery != nil {
		logData["error_query"] = errQuery.Error()
		repo.logger.
			WithFields(logData).
			WithError(errQuery).
			Errorln("error on insert query")
		return domain.CreateNewUserAccountResult{}, errQuery
	}

	idResult, errIdResult := execContextResp.LastInsertId()
	if errIdResult != nil {
		logData["error_id_result"] = errIdResult.Error()
		repo.logger.
			WithFields(logData).
			WithError(errQuery).
			Errorln("error on insert query")
		return domain.CreateNewUserAccountResult{}, errQuery
	}

	response := domain.CreateNewUserAccountResult{
		Id: idResult,
	}

	return response, nil
}

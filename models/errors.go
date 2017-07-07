package models

import "reflect"

func getTypeData(i interface{}) string {

	v := reflect.ValueOf(i)
	v = reflect.Indirect(v)
	return v.Type().String()
}

func isCustomeEror(err error) bool {

	switch getTypeData(err) {
	case "ErrorInternalServer":
		return true
	}
	return false
}

type ErrorInternalServer struct {
	Message string
}

func (e *ErrorInternalServer) Error() string {
	return e.Message
}

func NewErrorInternalServer() error {

	e := &ErrorInternalServer{
		Message: "Internal Server Error",
	}
	return e
}

type ErrorNotFound struct {
	Message string
}

func (e *ErrorNotFound) Error() string {
	return e.Message
}

func NewErrorNotFound() error {

	e := &ErrorNotFound{
		Message: "Your requsted item  is Not Found",
	}
	return e
}

type ErrorConflict struct {
	Message string
}

func (e *ErrorConflict) Error() string {
	return e.Message
}

func NewErrorConflict() error {

	e := &ErrorConflict{
		Message: "Your requsted item already Exist. Conflict!",
	}
	return e
}

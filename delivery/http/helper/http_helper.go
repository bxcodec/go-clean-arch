package http

import (
	"net/http"
	"reflect"
)

type HttpHelper struct {
}

func (u *HttpHelper) getTypeData(i interface{}) string {

	v := reflect.ValueOf(i)
	v = reflect.Indirect(v)
	return v.Type().String()
}

func (u *HttpHelper) GetStatusCode(err error) int {
	statusCode := http.StatusOK
	if err != nil {
		switch u.getTypeData(err) {
		case "models.ErrorUnauthorized":
			statusCode = http.StatusUnauthorized
		case "models.ErrorNotFound":
			statusCode = http.StatusNotFound
		case "models.ErrorConflict":
			statusCode = http.StatusConflict
		case "models.ErrorInternalServer":
			statusCode = http.StatusInternalServerError
		default:
			statusCode = http.StatusInternalServerError
		}
	}
	return statusCode
}

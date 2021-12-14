package usecases

import "net/http"

type JSONResponse interface {
	HTTPStatus(w http.ResponseWriter, code int, msg []byte)
}

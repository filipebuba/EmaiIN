package endpoints

import (
	"emailn/internal/internalError"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

type EndpoimFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointfunc EndpoimFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointfunc(w, r)
		if err != nil {

			if errors.Is(err, internalError.ErrInternal) {
				render.Status(r, 500)
			} else {
				http.Error(w, err.Error(), 400)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)

		}
	})
}

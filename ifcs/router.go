package ifcs

import "net/http"

type IRouter interface {
	PathPrefix(string) IRouter
	Handler(http.Handler) IRouter
	ServeHTTP(http.ResponseWriter, *http.Request)
	Use(...func(http.Handler) http.Handler)
}

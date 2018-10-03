package middleware

import "net/http"

type ContentTypeJson struct{}

// Logger provides logs for the accesses to the go server using the routes in routes.go
func (_ ContentTypeJson) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	next(rw, r)
}

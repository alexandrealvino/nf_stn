package middleware

import (
	"net/http"
)

// JoinMiddleWares joins authenticator and logger middleWares
func JoinMiddleWares(next http.HandlerFunc) http.Handler {
	return Logger(TokenAuthMiddleware(next))
}
//
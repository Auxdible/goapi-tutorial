package middleware

import (
	"errors"
	"net/http"

	"github.com/Auxdible/goapi-tutorial/api"
	"github.com/Auxdible/goapi-tutorial/internal/tools"
	log "github.com/sirupsen/logrus"
)

var errorUnauthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(errorUnauthorized)
			api.RequestErrorHandler(w, errorUnauthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(errorUnauthorized)
			api.RequestErrorHandler(w, errorUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

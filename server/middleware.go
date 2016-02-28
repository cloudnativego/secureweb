package server

import (
	"net/http"

	"github.com/astaxie/beego/session"
	"github.com/codegangsta/negroni"
)

func isAuthenticated(sessionManager *session.Manager) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		session, _ := sessionManager.SessionStart(w, r)
		defer session.SessionRelease(w)
		if session.Get("profile") == nil {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		} else {
			next(w, r)
		}
	}
}

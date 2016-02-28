package server

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/session"
)

func userHandler(sessionManager *session.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := sessionManager.SessionStart(w, r)
		defer session.SessionRelease(w)

		// Getting the profile from the session
		profile := session.Get("profile")

		//TODO: Do something compelling with profile data (html template?)
		fmt.Fprintf(w, "USER DATA: %+v", profile)
	}
}

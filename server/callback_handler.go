package server

import (
	// Don't forget this first import or nothing will work
	_ "crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/session"

	"golang.org/x/oauth2"
)

func callbackHandler(sessionManager *session.Manager, config *authConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Instantiating the OAuth2 package to exchange the Code for a Token
		conf := &oauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			RedirectURL:  config.CallbackURL,
			Scopes:       []string{"openid", "name", "email", "picture"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://" + config.Domain + "/authorize",
				TokenURL: "https://" + config.Domain + "/oauth/token",
			},
		}

		// Getting the Code that we got from Auth0
		code := r.URL.Query().Get("code")

		// Exchanging the code for a token
		token, err := conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Getting now the User information
		client := conf.Client(oauth2.NoContext, token)
		resp, err := client.Get("https://" + config.Domain + "/userinfo")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Reading the body
		raw, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshalling the JSON of the Profile
		var profile map[string]interface{}
		if err := json.Unmarshal(raw, &profile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, _ := sessionManager.SessionStart(w, r)
		defer session.SessionRelease(w)

		session.Set("id_token", token.Extra("id_token"))
		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)

		// Redirect to logged in page
		http.Redirect(w, r, "/user", http.StatusMovedPermanently)

	}
}

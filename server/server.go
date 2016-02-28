package server

import (
	"net/http"

	"github.com/astaxie/beego/session"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/cf-tools"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type authConfig struct {
	ClientID     string
	ClientSecret string
	Domain       string
	CallbackURL  string
}

//NewServer configures and returns a Negroni server
func NewServer(appEnv *cfenv.App) *negroni.Negroni {
	// HACK handle these failures for realzies
	authClientID, _ := cftools.GetVCAPServiceProperty("authzero", "id", appEnv)
	authSecret, _ := cftools.GetVCAPServiceProperty("authzero", "secret", appEnv)
	authDomain, _ := cftools.GetVCAPServiceProperty("authzero", "domain", appEnv)
	authCallback, _ := cftools.GetVCAPServiceProperty("authzero", "callback", appEnv)

	config := &authConfig{
		ClientID:     authClientID,
		ClientSecret: authSecret,
		Domain:       authDomain,
		CallbackURL:  authCallback,
	}

	//TODO: Create a code example using redis as a 'pro tip' for end of chapter.
	sessionManager, _ := session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go sessionManager.GC()

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, sessionManager, config)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, sessionManager *session.Manager, config *authConfig) {
	mx.HandleFunc("/", homeHandler(config))
	mx.HandleFunc("/callback", callbackHandler(sessionManager, config))
	mx.Handle("/user", negroni.New(
		negroni.HandlerFunc(isAuthenticated(sessionManager)),
		negroni.Wrap(http.HandlerFunc(userHandler(sessionManager))),
	))
	mx.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
}

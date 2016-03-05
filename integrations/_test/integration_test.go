package main

import (
	"fmt"
	"testing"

	"github.com/cloudfoundry-community/go-cfenv"
	. "github.com/cloudnativego/secureweb/server"
	"github.com/codegangsta/negroni"
)

var (
	server *negroni.Negroni
)

//NOTE: Wercker will need to have the AUTHZERO environment variables injected for local builds
func TestIntegration(t *testing.T) {
	fmt.Println("food")
	appEnv, _ := cfenv.Current()
	server = NewServer(appEnv)
	if server == nil {
		t.Error("Server should not be nil.")
	}
}

package main

import (
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/secureweb/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, _ := cfenv.Current()
	s := server.NewServer(appEnv)
	s.Run(":" + port)
}

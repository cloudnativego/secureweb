package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/secureweb/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Printf("FATAL: Could not retrieve CF environment: %v", err)
		os.Exit(1)
	}
	s := server.NewServer(appEnv)
	s.Run(":" + port)
}

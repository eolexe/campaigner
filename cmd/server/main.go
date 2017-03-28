package main

import (
	"flag"

	"log"

	"github.com/eolexe/campaigner/server"
	"github.com/eolexe/campaigner/shared/config"
	"github.com/eolexe/campaigner/shared/env"
)

var Version string = `not configured version. Please build binary with flags '-ldflags "-X main.Version=your_version"'`

func main() {
	configFilePath := flag.String("conf", "config/local.json", "Env config file")

	config := config.MustNewConfig(*configFilePath)
	e := env.MustNewEnvironment(config)
	e.Version = Version

	log.Print("🚀  starting env: " + config.Name)
	s := server.NewServer(e)
	log.Print("🚀  running server at: " + config.Server.String())
	err := s.Run()

	if err != nil {
		log.Panicf("Failed to start server: %s", err.Error())
	}
}

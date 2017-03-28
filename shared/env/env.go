package env

import (
	"github.com/eolexe/campaigner/shared/config"
	"github.com/satori/go.uuid"
)

type Environment struct {
	Version     string
	Config      config.Config
	UserCounter int64
	IsDebug     bool
}

func (e *Environment) PopUserCounter() {
	e.UserCounter = e.UserCounter + 1
}

func MustNewEnvironment(config config.Config) *Environment {
	e := &Environment{}
	e.Config = config

	return e
}

func NewMockEnvironment() *Environment {
	e := &Environment{
		Version: "i_am_stub_" + uuid.NewV4().String(),
		Config:  config.Config{},
	}

	return e
}

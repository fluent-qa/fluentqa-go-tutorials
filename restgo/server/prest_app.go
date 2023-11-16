package server

import (
	"github.com/prest/prest/cmd"
	"github.com/prest/prest/config"
)

func RunPrestServer() {
	config.Load()
	cmd.Execute()
}

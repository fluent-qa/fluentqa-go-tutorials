package main

import (
	"log/slog"
	"simple-server/server"
)

func main() {
	slog.Info("start Auth Server")
	server.RunRestGoAuthSqliteServer(9090)
}

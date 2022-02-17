package main

import (
	"knocker/pkg/handler"
	"knocker/pkg/tools"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	tools.InitLogger(6)
	tools.Logger.Debug("logger initialized")
	handler.HandleRequest()
}

package main

import (
	config "knocker/configs"
	"knocker/pkg/handler"
	"knocker/pkg/tools"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.InitConfig()
	tools.Logger.Debug(tools.InitLogger(6))
	handler.HandleRequest()
}

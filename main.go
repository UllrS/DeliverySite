package main

import (
	"knocker/pkg/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	handler.HandleRequest()
}

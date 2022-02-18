package repository

import (
	"database/sql"
	config "knocker/configs"
	"knocker/pkg/tools"
)

func repository_connect() (*sql.DB, error) {

	db, err := sql.Open(config.Config.GetSqlSettings())
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, err
	}
	return db, err
}

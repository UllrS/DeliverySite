package repository

import (
	"database/sql"
	"fmt"
	"knocker/models"
	"knocker/pkg/tools"
)

func Admin_AUTH(login string, pwd string) (bool, error) {

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return false, err
	}
	defer db.Close()
	var query_string = fmt.Sprintf("SELECT `token` FROM `admins` WHERE `login` = '%s' AND `password` = '%s'", login, pwd)

	res := db.QueryRow(query_string)
	var token string
	err = res.Scan(&token)
	if err != nil {
		tools.Logger.Error(err.Error())
		return false, err
	}
	return err == nil, err
}

func Select_admin_list() ([]models.SuperUser, error) {
	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprint("SELECT * FROM `admins`")
	res, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	var admin_list = []models.SuperUser{}
	for res.Next() {
		var superuser models.SuperUser
		err = res.Scan(&superuser.Id, &superuser.Login, &superuser.Password, &superuser.Token)
		if err != nil {
			tools.Logger.Warn(err.Error())
			return nil, err
		}

		admin_list = append(admin_list, superuser)
	}
	return admin_list, err
}
func Select_User_List() ([]models.User, error) {
	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprint("SELECT * FROM `users`")
	res, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	var user_list = []models.User{}
	for res.Next() {
		var user models.User
		err = res.Scan(&user.Id, &user.Name, &user.Password, &user.Token, &user.Basket)
		if err != nil {
			tools.Logger.Warn(err.Error())
			return nil, err
		}

		user_list = append(user_list, user)
	}
	return user_list, err
}
func Select_User(table_name string, login string, password string) (*sql.Rows, error) {

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `%s` WHERE `login` LIKE '%s' AND `password` LIKE '%s' LIMIT 1", table_name, login, password)
	res, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	return res, err
}

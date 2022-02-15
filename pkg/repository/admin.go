package repository

import (
	"fmt"
	"log"
)

func Admin_AUTH(login string, pwd string) (bool, error) {

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	defer db.Close()
	log.Println("DB admin auth")
	log.Println(login)
	log.Println(pwd)
	//var pwd_hash = sha256.Sum256([]byte(pwd)) //hex.EncodeToString()
	var query_string = fmt.Sprintf("SELECT `token` FROM `admins` WHERE `login` = '%s' AND `password` = '%s'", login, pwd)
	//
	res := db.QueryRow(query_string)
	var token string
	err = res.Scan(&token)
	if err != nil {
		return false, err
	}
	return err == nil, err
}

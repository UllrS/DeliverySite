package repository

import (
	"database/sql"
	"fmt"
	"knocker/models"
	"knocker/pkg/tools"
)

func Basket_update(user_id int, basket string) error {

	tools.Logger.Trace("start function")
	var query_string = fmt.Sprintf("UPDATE `users` SET `basket`='%s' WHERE `id` = '%d'", basket, user_id)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return err
	}
	defer db.Close()
	insert, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	defer insert.Close()
	tools.Logger.Trace("end function")
	return err
}
func Basket_load(field string, basket_map map[int]int) ([]models.Product, float32, error) {
	tools.Logger.Trace("start function")

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, 0, err
	}
	defer db.Close()
	product_list = []models.Product{}
	var summ float32
	for key, value := range basket_map {
		var product models.Product
		var query_string = fmt.Sprintf("SELECT * FROM `products` WHERE `%s` LIKE '%d'", field, key)
		res := db.QueryRow(query_string)
		err = res.Scan(&product.Id, &product.Name, &product.Price, &product.Category, &product.Anons, &product.Merch, &product.Portion, &product.Unit)

		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, 0, err
		}
		product.Qty = value
		product.Init_sum()
		summ += product.Sumprice
		product_list = append(product_list, product)
	}
	tools.Logger.Trace("end function")
	return product_list, summ, nil
}

func Insert_order(j models.Order) error {

	tools.Logger.Trace("start function")
	var query_string = fmt.Sprintf("INSERT INTO `orders`(`user`, `tel`, `shipping`, `basket`, `date`) VALUES ('%d', '%s', '%s', '%s', '%s')", j.User, j.Tel, j.Shipping, j.Basket, j.Date)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	defer db.Close()
	insert, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	defer insert.Close()
	tools.Logger.Trace("end function")
	return nil

}

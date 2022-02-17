package repository

import (
	"fmt"
	"knocker/models"
	"knocker/pkg/tools"
)

func Select_order_list() ([]models.Order, error) {
	tools.Logger.Trace("Start function")
	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprint("SELECT * FROM `orders` WHERE `status` LIKE '%'")
	res, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	var order_list = []models.Order{}
	for res.Next() {
		var order models.Order
		err = res.Scan(&order.Id, &order.Token, &order.User, &order.Shipping, &order.Basket, &order.Tel, &order.Status, &order.Date)
		if err != nil {
			tools.Logger.Warn(err.Error())
			return nil, err
		}

		order_list = append(order_list, order)
	}
	tools.Logger.Trace("end function")
	return order_list, err
}
func Delete_Order(id int) error {

	tools.Logger.Trace("start function")
	var query_string = fmt.Sprintf("DELETE FROM `orders` WHERE `id` = '%d'", id)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	tools.Logger.Trace("end function")
	return nil
}

func Status_Order_up(order_id int) error {

	tools.Logger.Trace("start function")
	tools.Logger.Tracef("order_id: ", order_id)
	var query_string = fmt.Sprintf("UPDATE `orders` SET `status`= '%d' WHERE `id` = '%d'", 1, order_id)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	tools.Logger.Trace("end function")
	return nil
}

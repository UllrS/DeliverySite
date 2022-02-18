package repository

import (
	"database/sql"
	"fmt"
	"knocker/models"
	"knocker/pkg/tools"
)

var merch_list = []models.Merchant{}
var product_list = []models.Product{}
var product_detail = models.Product{}
var order_detail = models.Order{}

func Select_all(table_name string) (*sql.Rows, error) {

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, err
	}
	defer db.Close()
	var query_string = fmt.Sprintf("SELECT * FROM `%s`", table_name)
	res, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	return res, err
}
func Select_merch_list_where(field string, field_value string) ([]models.Merchant, error) {

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `merchants` WHERE `%s` LIKE '%s'", field, field_value)
	res, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	merch_list = []models.Merchant{}
	for res.Next() {
		var merch models.Merchant
		err = res.Scan(&merch.Id, &merch.Name, &merch.Type, &merch.Addr, &merch.Anons, &merch.Img, &merch.Date)
		if err != nil {
			tools.Logger.Warn(err.Error())
			return nil, err
		}

		merch_list = append(merch_list, merch)
	}
	return merch_list, err
}
func Select_merch_where(merch_id int) (*models.Merchant, error) {

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `merchants` WHERE `id` = '%d'", merch_id)
	var merchant models.Merchant
	res := db.QueryRow(query_string)
	err = res.Scan(&merchant.Id, &merchant.Name, &merchant.Type, &merchant.Addr, &merchant.Anons, &merchant.Img, &merchant.Date)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	return &merchant, err
}
func Select_product_where(key string, value string) (*models.Product, error) {

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `products` WHERE `%s` = '%s'", key, value)
	var product models.Product
	res := db.QueryRow(query_string)
	err = res.Scan(&product.Id, &product.Name, &product.Price, &product.Category, &product.Anons, &product.Merch, &product.Portion, &product.Unit)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	return &product, err
}
func Select_product_list_where(key string, value string) ([]models.Product, error) {
	tools.Logger.Trace("start function")

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `products` WHERE `%s` = '%s'", key, value)

	res, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return nil, err
	}
	product_list = []models.Product{}
	for res.Next() {
		var product models.Product
		err = res.Scan(&product.Id, &product.Name, &product.Price, &product.Category, &product.Anons, &product.Merch, &product.Portion, &product.Unit)
		if err != nil {
			tools.Logger.Warn(err.Error())
			return nil, err
		}

		product_list = append(product_list, product)
	}
	tools.Logger.Trace("end function")
	return product_list, err
}
func Insert_product(j models.Product) int {

	tools.Logger.Trace("start function")
	var id int
	var query_string = fmt.Sprintf("INSERT INTO `products`(`name`, `price`, `category`, `anons`, `merch`, `portion`, `unit`) VALUES ('%s', '%.2f', '%s', '%s', '%d', '%.f', '%s')", j.Name, j.Price, j.Category, j.Anons, j.Merch, j.Portion, j.Unit)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Error(err.Error())
		return 0
	}
	defer db.Close()
	insert, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return 0
	}
	defer insert.Close()

	query_string = fmt.Sprintf("SELECT `id` FROM `products` WHERE `name` LIKE '%s' AND `category` LIKE '%s' AND `anons` LIKE '%s' AND `merch` LIKE '%d' AND `portion` LIKE '%.f' AND `unit` LIKE '%s'", j.Name, j.Category, j.Anons, j.Merch, j.Portion, j.Unit)
	res := db.QueryRow(query_string)
	err = res.Scan(&id)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return 0
	}

	tools.Logger.Trace("end function")
	return id
}
func Insert_merchant(j models.Merchant) int {

	tools.Logger.Trace("start function")
	var id int
	var query_string = fmt.Sprintf("INSERT INTO `merchants`(`name`, `type`, `addr`, `anons`, `date`) VALUES ('%s', '%s', '%s', '%s', '%s')", j.Name, j.Type, j.Addr, j.Anons, j.Date)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Warn(err.Error())
		return 0
	}
	defer db.Close()
	insert, err := db.Query(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return 0
	}
	defer insert.Close()

	query_string = fmt.Sprintf("SELECT `id` FROM `merchants` WHERE `name` LIKE '%s' AND `type` LIKE '%s' AND `addr` LIKE '%s' AND `anons` LIKE '%s' AND `date` LIKE '%s' LIMIT 1", j.Name, j.Type, j.Addr, j.Anons, j.Date)
	res := db.QueryRow(query_string)
	err = res.Scan(&id)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return 0
	}
	tools.Logger.Trace("end function")
	return id
}
func Update_merchant(j models.Merchant) error {

	tools.Logger.Trace("start function")
	var query_string = fmt.Sprintf("UPDATE `merchants` SET `name`='%s',  `addr`='%s', `anons`='%s', `date`='%s' WHERE `id`='%d'", j.Name, j.Addr, j.Anons, j.Date, j.Id)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	defer db.Close()

	if _, err = db.Exec(query_string); err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	tools.Logger.Trace("end function")
	return err
}
func Update_Product(j models.Product) error {

	tools.Logger.Trace("start function")
	var query_string = fmt.Sprintf("UPDATE `products` SET `name`='%s', `price`='%.2f', `category`='%s', `anons`='%s', `merch`='%d', `portion`='%.f', `unit`='%s' WHERE `id`='%d'", j.Name, j.Price, j.Category, j.Anons, j.Merch, j.Portion, j.Unit, j.Id)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	defer db.Close()

	if _, err = db.Exec(query_string); err != nil {
		tools.Logger.Error(err.Error())
		return err
	}
	tools.Logger.Trace("end function")
	return err
}

func Delete_merch(id int) error {

	tools.Logger.Trace("start function")
	var query_string = fmt.Sprintf("DELETE FROM `merchants` WHERE `id` = '%d'", id)

	db, err := repository_connect()
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	defer db.Close()
	// Delete Merch from DB
	_, err = db.Exec(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	// Delete all Merch products from DB
	query_string = fmt.Sprintf("DELETE FROM `products` WHERE `merch` = '%d'", id)

	_, err = db.Exec(query_string)
	if err != nil {
		tools.Logger.Warn(err.Error())
		return err
	}
	tools.Logger.Trace("end function")
	return nil

}
func Delete_prod(id int) error {

	tools.Logger.Trace("start function")
	var query_string = fmt.Sprintf("DELETE FROM `products` WHERE `id` = '%d'", id)

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

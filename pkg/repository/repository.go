package repository

import (
	"database/sql"
	"fmt"
	"knocker/models"
	"log"
)

var merch_list = []models.Merchant{}
var product_list = []models.Product{}
var product_detail = models.Product{}
var order_detail = models.Order{}

func repository_connect() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return db, err
}
func Select_all(table_name string) (*sql.Rows, error) {

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()
	var query_string = fmt.Sprintf("SELECT * FROM `%s`", table_name)
	res, err := db.Query(query_string)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return res, err
}
func Select_merch_list_where(field string, field_value string) ([]models.Merchant, error) {

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `merchants` WHERE `%s` LIKE '%s'", field, field_value)
	res, err := db.Query(query_string)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	merch_list = []models.Merchant{}
	for res.Next() {
		var merch models.Merchant
		err = res.Scan(&merch.Id, &merch.Name, &merch.Type, &merch.Addr, &merch.Anons, &merch.Img, &merch.Date)
		if err != nil {
			return nil, err
		}

		merch_list = append(merch_list, merch)
	}
	return merch_list, err
}
func Select_product_where(field string, field_value string) (*models.Product, error) {

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `products` WHERE `%s` = '%s'", field, field_value)
	var product models.Product
	res := db.QueryRow(query_string)
	err = res.Scan(&product.Id, &product.Name, &product.Price, &product.Category, &product.Anons, &product.Merch, &product.Portion, &product.Unit)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &product, err
}
func Select_product_list_where(field string, field_value string) ([]models.Product, error) {

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `products` WHERE `%s` = '%s'", field, field_value)

	res, err := db.Query(query_string)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	product_list = []models.Product{}
	for res.Next() {
		var product models.Product
		err = res.Scan(&product.Id, &product.Name, &product.Price, &product.Category, &product.Anons, &product.Merch, &product.Portion, &product.Unit)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		product_list = append(product_list, product)
	}
	return product_list, err
}
func Insert_product(j models.Product) int {

	var id int
	var query_string = fmt.Sprintf("INSERT INTO `products`(`name`, `price`, `category`, `anons`, `merch`, `portion`, `unit`) VALUES ('%s', '%.2f', '%s', '%s', '%d', '%.f', '%s')", j.Name, j.Price, j.Category, j.Anons, j.Merch, j.Portion, j.Unit)

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
	insert, err := db.Query(query_string)
	if err != nil {
		log.Println(err.Error())
	}
	defer insert.Close()

	fmt.Println(j)
	query_string = fmt.Sprintf("SELECT `id` FROM `products` WHERE `name` LIKE '%s' AND `category` LIKE '%s' AND `anons` LIKE '%s' AND `merch` LIKE '%d' AND `portion` LIKE '%.f' AND `unit` LIKE '%s'", j.Name, j.Category, j.Anons, j.Merch, j.Portion, j.Unit)
	res := db.QueryRow(query_string)
	err = res.Scan(&id)
	fmt.Println(res)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	fmt.Println("DB INSERT")
	fmt.Println(id)
	return id
}
func Insert_merchant(j models.Merchant) int {

	var id int
	var query_string = fmt.Sprintf("INSERT INTO `merchants`(`name`, `type`, `addr`, `anons`, `date`) VALUES ('%s', '%s', '%s', '%s', '%s')", j.Name, j.Type, j.Addr, j.Anons, j.Date)

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	defer db.Close()
	insert, err := db.Query(query_string)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	defer insert.Close()

	query_string = fmt.Sprintf("SELECT `id` FROM `merchants` WHERE `name` LIKE '%s' AND `type` LIKE '%s' AND `addr` LIKE '%s' AND `anons` LIKE '%s' AND `date` LIKE '%s' LIMIT 1", j.Name, j.Type, j.Addr, j.Anons, j.Date)
	res := db.QueryRow(query_string)
	err = res.Scan(&id)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	fmt.Println("DB INSERT")
	fmt.Println(id)
	return id
}
func Insert_order(j models.Order) {

	var query_string = fmt.Sprintf("INSERT INTO `orders`(`user`, `tel`, `billing`, `basket`) VALUES ('%d', '%s', '%s', '%s')", j.User, j.Tel, j.Billing, j.Basket)

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()
	insert, err := db.Query(query_string)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer insert.Close()

}
func Select_User(table_name string, login string, password string) (*sql.Rows, error) {

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db.Close()

	var query_string = fmt.Sprintf("SELECT * FROM `%s` WHERE `login` LIKE '%s' AND `password` LIKE '%s' LIMIT 1", table_name, login, password)
	res, err := db.Query(query_string)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return res, err
}

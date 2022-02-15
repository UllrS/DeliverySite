package repository

import (
	"database/sql"
	"fmt"
	"knocker/models"
	"log"
)

func Basket_update(user_id int, basket string) error {

	var query_string = fmt.Sprintf("UPDATE `users` SET `basket`='%s' WHERE `id` = '%d'", basket, user_id)

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer db.Close()
	insert, err := db.Query(query_string)
	if err != nil {
		return err
	}
	defer insert.Close()
	return err
}
func Basket_load(field string, basket_map map[int]int) ([]models.Product, float32, error) {

	db, err := repository_connect()
	if err != nil {
		log.Println(err.Error())
		return nil, 0, err
	}
	defer db.Close()
	fmt.Println("basket_map")
	fmt.Println(basket_map)
	product_list = []models.Product{}
	var summ float32
	for key, value := range basket_map {
		fmt.Println(key)
		fmt.Println(value)
		var product models.Product
		var query_string = fmt.Sprintf("SELECT * FROM `products` WHERE `%s` LIKE '%d'", field, key)
		res := db.QueryRow(query_string)
		fmt.Println(res)
		err = res.Scan(&product.Id, &product.Name, &product.Price, &product.Category, &product.Anons, &product.Merch, &product.Portion, &product.Unit)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, 0, err
		}
		summ += product.Price
		fmt.Println(value)
		product_list = append(product_list, product)
		// 	return nil, err
		// }
		// return res, err
	}
	return product_list, summ, nil
}

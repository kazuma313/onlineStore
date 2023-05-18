package ordermodels

import (
	"github.com/kazuma313/onlineStore/config"
	"github.com/kazuma313/onlineStore/entities"
)

func Create(order entities.Oreder) bool {
	result, err := config.DB.Exec("INSERT INTO `orders` (`order_id`, `costumer_id`, `product_id`) VALUES (NULL, ?, ?)", order.Costumer.Costumer_id, order.Product.Product_id)

	if err != nil {
		panic(err)
	}

	lastInsert, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsert > 0

}
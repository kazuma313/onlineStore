package costumermodels

import (
	"github.com/kazuma313/onlineStore/config"
	"github.com/kazuma313/onlineStore/entities"
)

func Create(Costumer entities.Costumer) bool{
	result, err := config.DB.Exec("INSERT INTO `costumer`(`email`, `username`, `password`) VALUES (?,?,?)", Costumer.Email, Costumer.Username, Costumer.Password)

	if err != nil {
		panic(err)
	}

	lastInsert, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsert > 0

}
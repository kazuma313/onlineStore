package categorymodels

import (
	"github.com/kazuma313/onlineStore/config"
	"github.com/kazuma313/onlineStore/entities"
)

func GetData() []entities.Category{
	rows, err := config.DB.Query("SELECT * FROM `productcategory`")
	
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	
	var categories [] entities.Category

	for rows.Next(){
		var category entities.Category
		err := rows.Scan(&category.Category, &category.Category_id)
		if err != nil{
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func GetDataProductByCategory(categori_id int) [] entities.ProductCategory{
	row, err := config.DB.Query("SELECT product, price FROM `product` INNER JOIN productcategory ON product.category_id = productcategory.category_id WHERE productcategory.category_id = ?",(categori_id))
	if err != nil {
		panic(err)
	}
	defer row.Close()

	var products [] entities.ProductCategory

	for row.Next(){
		var product entities.ProductCategory
		err := row.Scan(&product.Product, &product.Price)

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products

}
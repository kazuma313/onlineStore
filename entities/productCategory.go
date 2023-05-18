package entities

type Category struct {
	Category_id int
	Category    string
}

type ProductCategory struct {
	Product_id  int
	Category_id int
	Product     string `json:"product"`
	Price       int    `json:"price"`
}
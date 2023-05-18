package entities

type Category struct {
	Category_id int
	Category    string
}

type ProductCategory struct {
	Product string `json:"product"`
	Price   int    `json:"price"`
}
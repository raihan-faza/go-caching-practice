package data

type Product struct {
	ProductID       int64  `json:"product_id"`
	ProducName      string `json:"product_name"`
	ProductCategory string `json:"product_category"`
	ProductPrice    int64  `json:"product_price"`
}

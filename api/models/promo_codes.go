package models

type Promo_code struct {
	Promo_id 				int     `json:"promo_id"`
	Promo_name 				string  `json:"promo_name"`
	Promo_discount 			float64 `json:"promo_discount"`
	Promo_discount_type 	string  `json:"promo_discount_type"`
	Promo_order_limit_price float64 `json:"promo_order_limit_price"`
}

type PromoPrimaryKey struct {
	Promo_id int `json:"promo_id"`
}

type CreatePromo struct {
	Promo_name 				string  `json:"promo_name"`
	Promo_discount 			float64 `json:"promo_discount"`
	Promo_discount_type 	string  `json:"promo_discount_type"`
	Promo_order_limit_price float64 `json:"promo_order_limit_price"`
}



type GetListPromoRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPromoResponse struct {
	Count    int        `json:"count"`
	Promo_codes []*Promo_code `json:"promo_codes"`
}
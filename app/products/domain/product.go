package domain

type Products struct {
	Id          int32    `json:"id"`
	Name        string   `json:"name"`
	CategoryId  int32    `json:"category_id"`
	BrandId     int32    `json:"brand_id"`
	Description string   `json:"description"`
	IsNew       bool     `json:"is_new"`
	IsHot       bool     `json:"is_hot"`
	OnSale      bool     `json:"on_sale"`
	Click       int32    `json:"click"`
	Sale        int32    `json:"sale"`
	Favorite    int32    `json:"favorite"`
	MarkPrice   float32  `json:"mark_price"`
	ShopPrice   float32  `json:"shop_price"`
	Picture     string   `json:"picture"`
	Images      []string `json:"images"`
}

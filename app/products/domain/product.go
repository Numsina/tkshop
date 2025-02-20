package domain

type Products struct {
	Id           int32    `json:"id"`
	Name         string   `json:"name"`
	CategoryId   int32    `json:"category_id"`
	CategoryName string   `json:"category_name"`
	BrandId      int32    `json:"brand_id"`
	BrandName    string   `json:"brand_name"`
	Description  string   `json:"description"`
	IsNew        bool     `json:"is_new"`
	IsHot        bool     `json:"is_hot"`
	OnSale       bool     `json:"on_sale"`
	Click        int32    `json:"click"`
	Sale         int32    `json:"sale"`
	Favorite     int32    `json:"favorite"`
	MarkPrice    float32  `json:"mark_price"`
	ShopPrice    float32  `json:"shop_price"`
	Picture      string   `json:"picture"`
	Images       []string `json:"images"`
}

type Brands struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Categorys struct {
	Id             int32      `json:"id"`
	Name           string     `json:"name"`
	Level          int        `json:"level"`
	ParentId       int32      `json:"parent_id"`
	ParentCategory *Categorys `json:"parent_category"`
	RootId         int32      `json:"root_id"`
}

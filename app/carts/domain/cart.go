package domain

type Carts struct {
	Id      int32 `json:"id"`
	GoodsID int32 `json:"goods_id"`
	Nums    int32 `json:"nums"`
	Checked bool  `json:"checked"`
}

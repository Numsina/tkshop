package domain

type Orders struct {
	Id      int32   `json:"id"`
	OrderSn string  `json:"order_sn"`
	PayType string  `json:"pay_type"`
	Status  int32   `json:"status"`
	PayTime int64   `json:"pay_time"`
	Amount  float32 `json:"amount"`
	Address string  `json:"address"`
	Phone   string  `json:"phone"`
}

type OrderGoods struct {
	Id      int32   `json:"id"`
	OrderSn string  `json:"order_sn"`
	GoodsId int32   `json:"goods_id"`
	Price   float32 `json:"price"`
	Nums    int32   `json:"nums"`
}

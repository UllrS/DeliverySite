package models

type User struct {
	Id       int32    `json:"id"`
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Token    string   `json:"token"`
	Basket   string   `json:"basket"`
	Tags     []string `json:"tags"`
}
type Product struct {
	Id       int32   `json:"id"`
	Merch    int32   `json:"merch"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Anons    string  `json:"anons"`
	Unit     string  `json:"unit"`
	Price    float32 `json:"price"`
	Portion  float32 `json:"portion"`
}
type Merchant struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Addr  string `json:"addr"`
	Anons string `json:"anons"`
	Img   string `json:"img"`
	Date  string `json:"date"`
}
type Basket_Unit struct {
	Id  int `json:"id"`
	Qty int `json:"qty"`
}
type Order struct {
	Id      int32  `json:"id"`
	User    int    `json:"user"`
	Tel     string `json:"tel"`
	Billing string `json:"billing"`
	Token   string `json:"token"`
	Basket  string `json:"basket"`
}

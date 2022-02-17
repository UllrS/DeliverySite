package models

type SuperUser struct {
	Id       int32  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
type User struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Basket   string `json:"basket"`
}
type Product struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Merch    int32   `json:"merch"`
	Category string  `json:"category"`
	Anons    string  `json:"anons"`
	Unit     string  `json:"unit"`
	Price    float32 `json:"price"`
	Portion  float32 `json:"portion"`
	Qty      int     `json:"qty"`
	Sumqty   float32 `json:"sumqty"`
	Sumprice float32 `json:"sumprice"`
}

func (p *Product) Init_sum() {
	p.Sumqty = p.Portion * float32(p.Qty)
	p.Sumprice = p.Price * float32(p.Qty)
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
	Id       int32  `json:"id"`
	User     int    `json:"user"`
	Tel      string `json:"tel"`
	Shipping string `json:"shipping"`
	Token    string `json:"token"`
	Basket   string `json:"basket"`
	Status   int    `json:"status"`
	Date     string `json:"date"`
}

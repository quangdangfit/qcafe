package models

type Order struct {
	Code             string      `json:"code" bson:"code"`
	Lines            []OrderLine `json:"lines" bson:"lines"`
	UID              User        `json:"uid" bson:"uid"`
	TotalListedPrice float64     `json:"total_listed_price" bson:"total_listed_price"`
	TotalPrice       float64     `json:"total_price" bson:"total_price"`
}

type OrderLine struct {
	Product     Product `json:"product" bson:"product"`
	Quantity    int     `json:"quantity" bson:"quantity"`
	ListedPrice float64 `json:"listed_price" bson:"listed_price"`
	Price       float64 `json:"price" bson:"price"`
}

func (o Order) calculateTotalListedPrice() {
	o.TotalListedPrice = 0
	for _, line := range o.Lines {
		o.TotalListedPrice += line.ListedPrice
	}
}

func (o Order) calculatePrice() {
	o.TotalPrice = 0
	for _, line := range o.Lines {
		o.TotalPrice += line.Price
	}
}

func (line OrderLine) calculateListedPrice() {
	line.Price = line.Product.ListedPrice * float64(line.Quantity)
}

func (line OrderLine) calculatePrice() {
	line.Price = line.Product.Price * float64(line.Quantity)
}

package models

type Product struct {
	Code        string          `json:"code" bson:"code"`
	Category    ProductCategory `json:"category" bson:"category"`
	Description string          `json:"description" bson:"description"`
	ListedPrice float64         `json:"listed_price" bson:"listed_price"`
	Price       float64         `json:"price" bson:"price"`
	Promotion   Promotion       `json:"promotion" bson:"promotion"`
}

func (p *Product) ComputeRelatedFields() {
	p.CalculatePrice()
}

func (p *Product) CalculatePrice() {
	p.Price = p.ListedPrice - p.ListedPrice*p.Promotion.Percent
}

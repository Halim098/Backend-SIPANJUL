package Model

type Checkout_Detail struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
}

type Checkout struct {
	Items []Checkout_Detail `json:"items"`
	TotalAmount int `json:"totalAmount"`
}
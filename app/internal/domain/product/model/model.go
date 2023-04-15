package model

type Product struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ImageID       string `json:"image_Id"`
	Price         string `json:"price"`
	CurrencyID    string `json:"currency_Id"`
	Rating        string `json:"rating"`
	CategoryID    string `json:"category_Id"`
	Specification string `json:"specification"`
	CreatedAt     string `json:"created_At"`
	UpdatedAt     string `json:"updated_At"`
}

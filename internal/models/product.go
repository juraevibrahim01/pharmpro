package models

type Product_Req_login struct {
	AccessToken      string  `json:"access_token"`
	Product_name     string  `json:"product_name"`
	Firm_name        string  `json:"firm_name"`
	Country_product  string  `json:"country_product"`
	Amount           float64 `json:"amount"`
	Amount_sale      float64 `json:"amount_sale"`
	Markup           float32 `json:"markup"`
	Quantity_in_pack int32   `json:"quantity_in_pack"`
	Unit_id          int8    `json:"unit_id"`
}

type Product_Res_login struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Error       string `json:"error"`
}

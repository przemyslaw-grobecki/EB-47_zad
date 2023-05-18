package Models

type NewProductBinding struct {
	Id       uint    `json:"Id" binding:"required"`
	Name     string  `json:"Name" binding:"required"`
	Category string  `json:"Category" binding:"required"`
	Price    float64 `json:"Price" binding:"required"`
	Quantity uint    `json:"Quantity" binding:"required"`
	ImageURL string  `json:"ImageURL" binding:"required"`
}

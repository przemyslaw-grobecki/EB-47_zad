package Models

type PatchProductBinding struct {
	Name     *string  `json:"Name"`
	Category *string  `json:"Category"`
	Price    *float64 `json:"Price"`
	Quantity *uint    `json:"Quantity"`
	ImageURL *string  `json:"ImageURL"`
}

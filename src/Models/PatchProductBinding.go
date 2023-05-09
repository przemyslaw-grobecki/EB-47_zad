package Models

type PatchProductBinding struct {
	Name     *string `json:"Name"`
	Category *string `json:"Category"`
	Price    *uint   `json:"Price"`
	Quantity *uint   `json:"Quantity"`
	ImageURL *string `json:"ImageURL"`
}

package AddProduct

type Request struct {
	Id       uint
	Name     string
	Category string
	Price    float64
	Quantity uint
	ImageURL string
}

func NewRequest(id uint, name string, category string, price float64, quantity uint, imageURL string) *Request {
	return &Request{Id: id, Name: name, Category: category, Price: price, Quantity: quantity, ImageURL: imageURL}
}

package GetBasket

type Request struct {
	Id uint
}

func NewRequest(id uint) *Request {
	return &Request{Id: id}
}

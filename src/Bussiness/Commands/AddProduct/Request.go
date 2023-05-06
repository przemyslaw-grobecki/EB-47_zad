package AddProduct

type Request struct {
	Id   uint
	Name string
}

func NewRequest(id uint, name string) *Request {
	return &Request{Id: id, Name: name}
}

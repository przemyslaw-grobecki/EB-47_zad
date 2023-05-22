package RegisterUser

type Request struct {
	Email    string
	Password string
}

func NewRequest(email string, password string) *Request {
	return &Request{Email: email, Password: password}
}

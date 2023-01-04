package logistic

type PostService struct {
	a addressService
}

func NewPostService(service addressService) PostService {
	return PostService{service}
}

func (p PostService) getPostCode(customerID int) string {
	address := p.a.GetAddress(customerID, WorkAddress)
	if address == nil {
		address = p.a.GetAddress(customerID, HomeAddress)
	}
	return address.PostCode
}

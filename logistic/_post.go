package logistic

type postService struct {
	a addressService
}

func NewPostService(service addressService) postService {
	return postService{service}
}

func (p postService) getPostCode(customerID int64) string {
	address := p.a.GetAddress(customerID, WorkAddress)
	if address == nil {
		address = p.a.GetAddress(customerID, HomeAddress)
	}
	return address.PostCode
}

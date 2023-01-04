package logistic

type PostService struct {
	l Locator
}

func NewPostService(locator Locator) PostService {
	return PostService{locator}
}

func (p PostService) GetPostCode(customerID int64) string {
	address := p.l.GetAddress(customerID, WorkAddress)
	if address == nil {
		address = p.l.GetAddress(customerID, HomeAddress)
	}
	return address.PostCode
}

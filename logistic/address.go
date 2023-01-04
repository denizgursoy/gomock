//go:generate mockgen -source=address.go -destination=mock_address.go -package=logistic
package logistic

const (
	HomeAddress = "Home"
	WorkAddress = "Work"
)

type Locator interface {
	GetAddress(customerID int64, addressType string) *Address
}

type Address struct {
	PostCode string
}

type addressService struct {
}

func (a addressService) GetAddress(customerID int64, addressType string) *Address {
	// Does its jobs and return the address
	// Go to DB to fetch realy data
	return &Address{}
}

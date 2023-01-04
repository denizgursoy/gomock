package logistic

type mockAddressService struct{}

func (a mockAddressService) GetAddress(customerID int, addressType string) *Address {
	if customerID == 1 {
		return &Address{"2544TT"}
	} else if customerID == 2 {
		return &Address{"1111TK"}
	} else if customerID == 3 {
		return &Address{"6789GH"}
	}
	return nil
}

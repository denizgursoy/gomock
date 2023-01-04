package logistic

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostService_GetPostCode(t *testing.T) {
	t.Run("should return work address if it is not nil", func(t *testing.T) {
		//service := NewPostService(PostService{})
		//code := service.GetPostCode(1)
		//
		//require.Equal(t, "2544TT", code)
	})

	t.Run("should return work address if it is not nil", func(t *testing.T) {
		service := NewPostService(mockAddressService{})
		code := service.GetPostCode(1)

		require.Equal(t, "2544TT", code)
	})

	t.Run("should return work address if it is not nil", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)
		mockLocator.EXPECT().GetAddress(int64(1), WorkAddress).Return(&Address{"2544TT"})

		service := NewPostService(mockLocator)

		code := service.GetPostCode(1)

		require.Equal(t, "2544TT", code)
	})

	t.Run("should return home address if work address is nil", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)
		mockLocator.EXPECT().GetAddress(int64(2), WorkAddress).Return(nil)
		mockLocator.EXPECT().GetAddress(int64(2), HomeAddress).Return(&Address{"1111TK"})

		service := NewPostService(mockLocator)

		code := service.GetPostCode(2)

		require.Equal(t, "1111TK", code)
	})

	t.Run("arguments matcher", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)

		// Eq makes sure that first argument is equal to value
		mockLocator.EXPECT().GetAddress(gomock.Eq(int64(2)), WorkAddress)
		// Not makes sure that first argument is not equal to value
		mockLocator.EXPECT().GetAddress(gomock.Not(int64(1)), WorkAddress)
		// Any is like a placeholder for the argument, it does not check value
		mockLocator.EXPECT().GetAddress(gomock.Any(), HomeAddress)
		// Nil makes sure that first argument is nil
		mockLocator.EXPECT().GetAddress(gomock.Nil(), HomeAddress)
		// Len method checks the length of the first argument is equal to value.
		// It can be useful for checking size of arrays and maps arguments
		mockLocator.EXPECT().GetAddress(gomock.Len(1), HomeAddress)
		// All makes sure that argument matched by all matchers provided
		mockLocator.EXPECT().GetAddress(gomock.All(gomock.Not(int64(1)), gomock.Not(int64(2))), HomeAddress)
		// AssignableToTypeOf checks if the types of the arguments matches
		mockLocator.EXPECT().GetAddress(gomock.AssignableToTypeOf(int64(1)), HomeAddress)

		mockLocator.EXPECT().GetAddress(gomock.Any(), gomock.Any()).
			DoAndReturn(func(customerID int64, addressType string) *Address {
				if customerID == 1 {
					return &Address{"2544TT"}
				} else if customerID == 2 {
					return &Address{"1111TK"}
				} else if customerID == 3 {
					return &Address{"6789GH"}
				}
				return nil
			})
	})
}

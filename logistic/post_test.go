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

	t.Run("arguments matcher", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)

		// Eq makes sure that first argument is equal to value
		mockLocator.EXPECT().GetAddress(gomock.Eq(2), WorkAddress).Return(nil)
		// Not makes sure that first argument is not equal to value
		mockLocator.EXPECT().GetAddress(gomock.Not(1), WorkAddress).Return(nil)
		// Any is like a placeholder for the argument, it does not check value
		mockLocator.EXPECT().GetAddress(gomock.Any(), HomeAddress).Return(nil)
		// Nil makes sure that first argument is nil
		mockLocator.EXPECT().GetAddress(gomock.Nil(), HomeAddress).Return(nil)
		// Len method checks the length of the first argument is equal to value.
		// It can be useful for checking size of arrays and maps arguments
		mockLocator.EXPECT().GetAddress(gomock.Len(1), HomeAddress).Return(nil)
		// All makes sure that argument matched by all matchers provided
		mockLocator.EXPECT().GetAddress(gomock.All(gomock.Not(1), gomock.Not(2)), HomeAddress).Return(nil)
		// AssignableToTypeOf checks if the types of the arguments matches
		mockLocator.EXPECT().GetAddress(gomock.AssignableToTypeOf(1), HomeAddress).Return(nil)
	})

	t.Run("custom implementation of return", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)
		mockLocator.
			EXPECT().
			GetAddress(gomock.Any(), gomock.Any()).
			DoAndReturn(func(customerID int, addressType string) *Address {
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

	t.Run("changing orders of the call", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)
		fooCall := mockLocator.EXPECT().GetAddress(gomock.Any(), gomock.Any()).Return(&Address{"2544TT"})
		barCall := mockLocator.EXPECT().GetAddress(gomock.Any(), gomock.Any()).Return(&Address{"1111TK"})

		// Put one call after another call
		fooCall.After(barCall)

		// Keep the order as they are in the arguments
		gomock.InOrder(barCall, fooCall)
	})

	t.Run("repeating the calls ", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)
		// Times repeat call n times
		// In this case, following call will/must be executed twice
		mockLocator.EXPECT().GetAddress(1, gomock.Any()).Return(&Address{"1111TK"}).Times(1)
		// For the rest of the call, following call will be executed
		// It is expected to be called 0 or more times
		mockLocator.EXPECT().GetAddress(2, gomock.Any()).Return(nil).AnyTimes()
		// It is expected to be called at most twice
		mockLocator.EXPECT().GetAddress(3, gomock.Any()).Return(nil).MaxTimes(2)
		// It is expected to be called at least once
		mockLocator.EXPECT().GetAddress(4, gomock.Any()).Return(nil).MinTimes(1)

		service := NewPostService(mockLocator)

		service.GetPostCode(2)
	})
}

func TestPostService_GetPostCode_WithMock(t *testing.T) {
	t.Run("should return work address if it is not nil", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)
		// if you inspect the mock, you can see that it has all the methods of the interface
		// In tests, your code will call the methods on the mocks
		// mockLocator.GetAddress(1,WorkAddress)

		mockLocator.EXPECT().GetAddress(1, WorkAddress).Return(&Address{"2544TT"})

		service := NewPostService(mockLocator)

		code := service.GetPostCode(1)

		require.Equal(t, "2544TT", code)
	})
}

func TestPostService_GetPostCode_TwoCalls(t *testing.T) {
	t.Run("should return home address if work address is nil", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)
		mockLocator.EXPECT().GetAddress(2, WorkAddress).Return(nil)
		mockLocator.EXPECT().GetAddress(2, HomeAddress).Return(&Address{"1111TK"})

		service := NewPostService(mockLocator)

		code := service.GetPostCode(2)

		require.Equal(t, "1111TK", code)
	})
}

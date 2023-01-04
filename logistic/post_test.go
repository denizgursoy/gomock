package logistic

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostService_GetPostCode(t *testing.T) {
	t.Run("should return work address if it exists", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockLocator := NewMockLocator(controller)

		mockLocator.EXPECT().GetAddress(1, WorkAddress).Return(&Address{"2544TT"})

		service := NewPostService(mockLocator)

		code := service.GetPostCode(1)

		require.Equal(t, "2544TT", code)
	})

	t.Run("should return work address if it exists", func(t *testing.T) {
		service := NewPostService(mockAddressService{})
		code := service.GetPostCode(1)

		require.Equal(t, "2544TT", code)
	})

	t.Run("should return work address if it exists", func(t *testing.T) {
		//service := NewPostService(PostService{})
		//code := service.GetPostCode(1)
		//
		//require.Equal(t, "2544TT", code)
	})
}

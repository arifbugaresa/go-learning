package main

import (
	"github.com/stretchr/testify/assert"
	"go-learning/modules/user"
	"testing"
)

func TestMockingService(t *testing.T) {
	t.Run("success mocking service", func(t *testing.T) {
		srvMock := &user.UserServiceMock{}

		srvMock.
			On("GetMessage").
			Return("Message Mocking", " Test")

		message := GetMessageEndpointExampleForUnitTest(srvMock)

		assert.Equal(t, "Message Mocking Test", message)
	})
}

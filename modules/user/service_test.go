package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Data     interface{}
	Expected interface{}
	Actual   interface{}
}

func TestUser(t *testing.T) {
	testCases := []TestCase{
		{
			Name:     "Greeting to Arif",
			Data:     "arif",
			Expected: "Hello arif",
		},
		{
			Name:     "Greeting to Bugaresa",
			Data:     "bugaresa",
			Expected: "Hello bugaresa",
		},
		{
			Name:     "Greeting to Bugaresa Fail",
			Data:     "bugaresa",
			Expected: "Hello bugaresa Fail",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			repoMock := UserRepositoryMock{}

			// set mock repository
			repoMock.On("GetMessage", "arifbugaresa").Return("Hello "+tc.Data.(string), nil)

			userSrv := NewUserService(&repoMock)

			message := userSrv.GetMessage()

			assert.Equal(t, tc.Expected, message)

		})
	}
}

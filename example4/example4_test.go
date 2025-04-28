package example4

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockDBClient struct {
	mock.Mock
}

func (m *MockDBClient) List() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestFunction1(t *testing.T) {
	var testCases = []struct {
		data string
		err  error
		want string
	}{
		{data: "a,b,c", err: nil, want: "a,b,c"},
		{data: "d,e,f", err: nil, want: "d,e,f"},
		{data: "d,e,f", err: errors.New("something bad happened"), want: "d,e,f"},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("data: %v, err: %v", testCase.data, testCase.err)
		t.Run(testName, func(t *testing.T) {
			//given
			client := new(MockDBClient)
			service := NewService(client)

			//when
			client.On("List").Return(testCase.want, testCase.err)
			got, err := service.Function1()

			//then
			client.AssertExpectations(t)
			assert.Equal(t, testCase.want, got)
			assert.Equal(t, testCase.err, err)
		})
	}
}

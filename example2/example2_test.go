package example2

import (
	"fmt"
	"testing"
)

func TestFunction1(t *testing.T) {
	var testCases = []struct {
		data []string
		want string
	}{
		{[]string{"a", "b", "c"}, "a,b,c"},
		{[]string{"d", "e", "f"}, "d,e,f"},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("data: %v", testCase.data)
		t.Run(testName, func(t *testing.T) {
			//given
			client := &DBClient{
				data: testCase.data,
			}

			//when
			got, err := Function1(client)

			//then
			if err != nil {
				t.Errorf("received an error: %v", err)
			}
			if got != testCase.want {
				t.Errorf("got %v want %v", got, testCase.want)
			}
		})
	}
}

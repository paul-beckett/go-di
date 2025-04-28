package example1

import "testing"

func TestFunction1(t *testing.T) {
	//given
	want := "a,b,c"

	//when
	got, err := Function1()

	//then
	if err != nil {
		t.Errorf("received an error: %v", err)
	}
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

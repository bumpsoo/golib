package err_test

import (
	"fmt"
	"testing"

	"github.com/bumpsoo/golib/err"
)

func ToStr(int) (string, error) {
	return "foo", fmt.Errorf("some error")
}

func AnotherErr(string) (string, error) {
	return "foo foo foo", fmt.Errorf("another error")
}

func TestWrap(t *testing.T) {
	integer := err.Wrap(3, nil)
	w := err.Bind(integer, ToStr)
	ww := err.Bind(w, AnotherErr)
	s, err := ww()
	if err != nil && err.Error() != "some error" {
		t.Fatal("error", err.Error())
	}
	if len(s) > 0 {
		t.Fatal("error", s)
	}
}

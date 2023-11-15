package err_test

import (
	"strconv"
	"testing"

	"github.com/bumpsoo/golib/err"
)

func TestWrap(t *testing.T) {
	v := 3
	wr := err.Wrap(v, nil)
	w := err.Do(&wr, func(i int) (string, error) {
		return strconv.Itoa(i), nil
	})
}

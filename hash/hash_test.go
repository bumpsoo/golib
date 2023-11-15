package hash_test

import (
	"testing"

	"github.com/bumpsoo/golib/hash"
)

func TestMap(t *testing.T) {
	target := map[int]string{
		3: "foo",
		5: "foo",
	}
	ret := hash.Map(target, func(s string) string {
		return s + "bar"
	})
	for _, val := range ret {
		if val != "foobar" {
			t.Fatal(val, "wrong")
		}
	}
}

// erroR_test.go
package erroR

import (
	"fmt"
	"testing"
)

func TestRnInfile(t *testing.T) {
	a := RnKnew("bB1", "SomeBug!")
	b := RnIwrap("bE2", a)
	c := RnIwrap("bG3", b)
	d := RnGwrapErrList("bI4", nil, c, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func TestRnOut(t *testing.T) {
	a := RnKnew("FbB1", "SomeBug!")
	b := RnIwrap("FbE2", a)
	c := RnKnew("FbI4", "err==nil, Bug Here!")

	d := RnGwrapErrList("bI4", nil, c, c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	d = RnGwrapErrList("bI4", nil, nil, nil)
	fmt.Println("nil: ", d)
}

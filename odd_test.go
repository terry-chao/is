package is

import "testing"

func TestOdd(t *testing.T) {
	ret := Odd(0)
	if ret != false {
		t.Error()
	}
	ret = Odd(1)
	if ret != true {
		t.Error()
	}
	ret = Odd(2)
	if ret != false {
		t.Error()
	}
}

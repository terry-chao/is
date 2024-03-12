package is

import "testing"

func TestEven(t *testing.T) {
	ret := Even(0)
	if ret != true {
		t.Error()
	}
	ret = Even(1)
	if ret != false {
		t.Error()
	}
	ret = Even(2)
	if ret != true {
		t.Error()
	}
}

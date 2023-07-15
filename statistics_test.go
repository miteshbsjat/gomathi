package gomathi

import (
	"testing"
)

func TestMean(t *testing.T) {
	ret := mean(1, 2, 3, 4)

	if ret != 2.50 {
		t.Fail()
	}
}

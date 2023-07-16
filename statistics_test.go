package gomathi

import (
	"testing"
)

func TestMean(t *testing.T) {
	ret := Mean(1, 2, 3, 4)
	if ret != 2.50 {
		t.Fail()
	}

	ret = Mean()
	if ret != 0.0 {
		t.Fail()
	}
}

func TestMedian(t *testing.T) {
	ret := Median(3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1, 1, 1, 111)
	if ret != 2 {
		t.Fail()
	}

	ret = Median()
	if ret != 0.0 {
		t.Fail()
	}

	ret = Median(1, 2, 3)
	if ret != 2 {
		t.Fail()
	}

	ret = Median(1, 2, 3, 4)
	if ret != 2.5 {
		t.Fail()
	}
}

func TestQuantil(t *testing.T) {
	ret := Quantile(0)
	if ret != 0.0 {
		t.Fail()
	}
	ret = Quantile(0.8, 3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1, 1, 1, 111)
	if ret != 3 {
		t.Errorf("Quantile ret = %f\n", ret)
	}
	ret = Quantile(0.9, 3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1, 1, 1, 111)
	if ret != 22 {
		t.Errorf("Quantile ret = %f\n", ret)
	}
}

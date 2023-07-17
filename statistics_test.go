package gomathi

import (
	"testing"
)

func TestMean(t *testing.T) {
	arr := make([]int, 4)
	for i := 0; i < 4; i++ {
		arr[i] = 4 - i
	}
	expected := 2.5
	ret := Mean(arr)
	if ret != 2.50 {
		t.Errorf("returned=%f != expected=%f", ret, expected)
	}

	expected = 0.0
	ret = Mean(make([]float64, 0))
	if ret != expected {
		t.Errorf("returned=%f != expected=%f", ret, expected)
	}

	expected = 0.3
	arrf := []float64{0.3, 0.1, 0.2, 0.4, 0.6, 0.2}
	before := arrf[0]
	ret = Mean(arrf)
	if ret != expected {
		t.Errorf("returned=%f != expected=%f", ret, expected)
	}
	after := arrf[0]
	if before != after {
		t.Errorf("before=%f != after=%f", before, after)
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

func TestQuantile(t *testing.T) {
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

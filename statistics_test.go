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
	arrf := []float64{3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1, 1, 1, 111}
	ret := Median(arrf)
	expected := 2.0
	if ret != expected {
		t.Errorf("returned=%f != expected=%f", ret, expected)
	}

	ret = Median(make([]int64, 0))
	if ret != 0.0 {
		t.Fail()
	}

	arri := []int{3, 3, 3, 3, 2, 22, 2, 2, 1, 2}
	before := arri[0]
	ret = Median(arri)
	expected = 2.0
	if ret != expected {
		t.Errorf("returned=%f != expected=%f", ret, expected)
	}
	after := arri[0]
	if before != after {
		t.Errorf("before=%d != after=%d", before, after)
	}
}

func TestQuantile(t *testing.T) {
	ret := Quantile(0, make([]int, 0))
	if ret != 0.0 {
		t.Fail()
	}
	ret = Quantile(0.8, []int{3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1, 1, 1, 111})
	if ret != 3 {
		t.Errorf("Quantile ret = %d\n", ret)
	}
	retf := Quantile(0.9, []float64{3, 3, 3, 3, 2, 22, 2, 2, 2, 2, 1, 1, 1, 1, 1, 111})
	if retf != 22 {
		t.Errorf("Quantile ret = %f\n", retf)
	}
}

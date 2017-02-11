package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{1, 2, 3},
	{4, 5, 9},
	{127, 1, 128},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{1, 2, 3},
	{-2, 5, 3},
	{2147483640, 8, 2147483648},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{1, 2, 3},
	{-2, 5, 3},
	{127, 1, 128},
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	{1, 2, 3},
	{-2, 5, 3},
	{1.987346743760946923759723957237590235, 3.8764564959549834643760436834634964, 5.8638032397159303881357676407010866344},
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%g, %g) returned %g, expected %g", v.n1, v.n2, val, v.expected)
		}
	}
}

var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{1, 2, 3},
	{-2, 5, 3},
	{9223372036854775800, 9, 9223372036854775809},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 2, want: 3},
		{a: -2, b: 2, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 12
	got := calculator.Multiply(3, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b, want  float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 1, b: 1, want: 1, errExpected: false},
		{a: 10, b: 2, want: 5, errExpected: false},
		{a: 6, b: 6, want: 1, errExpected: false},
		{a: 100, b: 0, want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if (tc.want != got) && err == nil {
			t.Fatalf("want %f, got %f. unexpected error status: nil", tc.want, got)
		}
		if tc.errExpected && err == nil {
			t.Fatalf("want %f, got %f. unexpected error status: nil", tc.want, got)
		}
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
		if (err != nil) && (!tc.errExpected) {
			t.Errorf("got an error %f", err)
		}
	}
}

// File: go/main_test.go
package go_calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2.5, 3.5)
	expected := 6.0
	if result != expected {
		t.Errorf("Add(2.5, 3.5) = %f; want %f", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := Sub(10.0, 4.0)
	expected := 6.0
	if result != expected {
		t.Errorf("Sub(10.0, 4.0) = %f; want %f", result, expected)
	}
}

func TestMul(t *testing.T) {
	result := Mul(7.0, 6.0)
	expected := 42.0
	if result != expected {
		t.Errorf("Mul(7.0, 6.0) = %f; want %f", result, expected)
	}
}

func TestDiv(t *testing.T) {
	result := Div(10.0, 2.0)
	expected := 5.0
	if result != expected {
		t.Errorf("Div(10.0, 2.0) = %f; want %f", result, expected)
	}
}

// This test checks that the Div function panics when dividing by zero.
func TestDivByZero(t *testing.T) {
	// defer is used to run code at the end of the function.
	// recover() catches the panic message.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Div() did not panic")
		}
	}()
	Div(10.0, 0.0) // This line should cause a panic
}

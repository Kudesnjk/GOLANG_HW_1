package main

import (
	"testing"
)

func TestAddSuccess(t *testing.T) {
	polishStr := "1 2 +"
	result, err := getResult(polishStr)
	if result != 3 && err != nil {
		t.Fatalf("Add test failed on %s", err)
	}
}

func TestSubSuccess(t *testing.T) {
	polishStr := "6 3 -"
	result, err := getResult(polishStr)
	if result != 3 && err != nil {
		t.Fatalf("Sub test failed on %s", err)
	}
}

func TestMulSuccess(t *testing.T) {
	polishStr := "3 2 *"
	result, err := getResult(polishStr)
	if result != 6 && err != nil {
		t.Fatalf("Mul test failed on %s", err)
	}
}

func TestDivSuccess(t *testing.T) {
	polishStr := "6 3 /"
	result, err := getResult(polishStr)
	if result != 2 && err != nil {
		t.Fatalf("Div test failed on %s", err)
	}
}

func TestDivFail(t *testing.T) {
	polishStr := "6 0 /"
	_, err := getResult(polishStr)
	if err == nil {
		t.Fatalf("Div test not failed on %s", err)
	}
}

func TestPolishNotationSuccess(t *testing.T) {
	input := "2 * ( 4 +6)"
	result, err := getPolishNotation(input)
	if result != "4 6 + 2 *" && err != nil {
		t.Fatalf("Polish notation test failed on %s, result is %s", result, err)
	}
}

func TestPolishNotationFail(t *testing.T) {
	input := "2 $$ ( 4 +6)"
	_, err := getPolishNotation(input)
	if err == nil {
		t.Fatalf("Polish notation test not failed on %s", input)
	}
}

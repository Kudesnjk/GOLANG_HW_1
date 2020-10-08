package main

import (
	"reflect"
	"testing"
)

func TestDefaultBehaivourSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	expected := []string{
		"I love music.",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check default behaviour failed")
	}
}

func TestDefaultBehaivourFail(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}
	var option options
	expected := []string{
		"I love music.",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check default behaviour not failed")
	}
}

func TestPrintEntriesCountBehaivourSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.PrintEntriesCount = true
	expected := []string{
		"3 I love music.",
		"2 I love music of Kartik.",
		"1 Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check PrintEntriesCount behaviour failed")
	}
}

func TestPrintEntriesCountBehaivourFail(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}
	var option options
	option.PrintEntriesCount = true
	expected := []string{
		"3 I love music.",
		"2 I love music of Kartik.",
		"1 Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check PrintEntriesCount behaviour not failed")
	}
}

func TestPrintRepeatedBehaivourSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.PrintRepeated = true
	expected := []string{
		"I love music.",
		"I love music of Kartik.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check PrintRepeated behaviour failed")
	}
}

func TestPrintRepeatedBehaivourFail(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.PrintRepeated = true
	expected := []string{
		"I love music.",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check PrintRepeated behaviour not failed")
	}
}

func TestPrintUnRepeatedBehaivourSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.PrintUnRepeated = true
	expected := []string{
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check PrintUnRepeated behaviour failed")
	}
}

func TestPrintUnRepeatedBehaivourFail(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.PrintUnRepeated = true
	expected := []string{
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check PrintUnRepeated behaviour not failed")
	}
}

func TestWithoutRegisterBehaivourSuccess(t *testing.T) {
	data := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"I love MuSIC of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.WithoutRegister = true
	expected := []string{
		"I LOVE MUSIC.",
		"I love MuSIC of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check WithoutRegister behaviour failed")
	}
}

func TestWithoutRegisterBehaivourFail(t *testing.T) {
	data := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"I love MuSIC of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.WithoutRegister = true
	expected := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I love MuSIC of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check WithoutRegister behaviour not failed")
	}
}

func TestSkipNumFieldsBehaivourSuccess(t *testing.T) {
	data := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.SkipNumFields = 1
	expected := []string{
		"We love music.",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check SkipNumFields behaviour failed")
	}
}

func TestSkipNumFieldsBehaivourFail(t *testing.T) {
	data := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.SkipNumFields = 1
	expected := []string{
		"I love music.",
		"We love music.",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check SkipNumFields behaviour not failed")
	}
}

func TestSkipNumCharsBehaivourSuccess(t *testing.T) {
	data := []string{
		"A love music.",
		"I love music.",
		"C love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.SkipNumChars = 1
	expected := []string{
		"A love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check SkipNumChars behaviour failed")
	}
}

func TestSkipNumCharsBehaivourFail(t *testing.T) {
	data := []string{
		"A love music.",
		"I love music.",
		"C love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.SkipNumChars = 1
	expected := []string{
		"A love music.",
		"I love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check SkipNumChars behaviour not failed")
	}
}

func TestCDUArgsToOptionSuccess(t *testing.T) {
	args := []string{
		"-c",
		"-d",
		"-u",
	}
	_, err := argsToOptions(args)
	if err == nil {
		t.Fatalf("C D U flags check failed")
	}
}

func TestCDUArgsToOptionFail(t *testing.T) {
	args := []string{
		"-c",
	}
	_, err := argsToOptions(args)
	if err != nil {
		t.Fatalf("C D U flags check not failed")
	}
}

func TestArgsToOptionSuccess(t *testing.T) {
	args := []string{
		"-c",
		"-f",
		"1",
		"-i",
	}
	expected := options{
		PrintEntriesCount: true,
		SkipNumFields:     1,
		WithoutRegister:   true,
	}
	result, err := argsToOptions(args)
	if err != nil && reflect.DeepEqual(expected, result) {
		t.Fatalf("Flags check failed")
	}
}

func TestArgsToOptionFail(t *testing.T) {
	args := []string{
		"-c",
		"-f",
		"-i",
		"1",
	}
	expected := options{
		PrintEntriesCount: true,
		SkipNumFields:     1,
		WithoutRegister:   true,
	}
	result, err := argsToOptions(args)
	if err == nil || reflect.DeepEqual(expected, result) {
		t.Fatalf("Flags check not failed")
	}
}

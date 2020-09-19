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

func TestCFlagBehaivourSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.cFlag = true
	expected := []string{
		"3 I love music.",
		"2 I love music of Kartik.",
		"1 Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check CFlag behaviour failed")
	}
}

func TestCFlagBehaivourFail(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}
	var option options
	option.cFlag = true
	expected := []string{
		"3 I love music.",
		"2 I love music of Kartik.",
		"1 Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check CFlag behaviour not failed")
	}
}

func TestDFlagBehaivourSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.dFlag = true
	expected := []string{
		"I love music.",
		"I love music of Kartik.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check DFlag behaviour failed")
	}
}

func TestDFlagBehaivourFail(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.dFlag = true
	expected := []string{
		"I love music.",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check DFlag behaviour not failed")
	}
}

func TestUFlagBehaivourSuccess(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.uFlag = true
	expected := []string{
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check UFlag behaviour failed")
	}
}

func TestUFlagBehaivourFail(t *testing.T) {
	data := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.uFlag = true
	expected := []string{
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check UFlag behaviour not failed")
	}
}

func TestIFlagBehaivourSuccess(t *testing.T) {
	data := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"I love MuSIC of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.iFlag = true
	expected := []string{
		"I LOVE MUSIC.",
		"I love MuSIC of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check iFlag behaviour failed")
	}
}

func TestIFlagBehaivourFail(t *testing.T) {
	data := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"I love MuSIC of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.iFlag = true
	expected := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I love MuSIC of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check iFlag behaviour not failed")
	}
}

func TestFFlagBehaivourSuccess(t *testing.T) {
	data := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.fFlag = 1
	expected := []string{
		"We love music.",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check fFlag behaviour failed")
	}
}

func TestFFlagBehaivourFail(t *testing.T) {
	data := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.fFlag = 1
	expected := []string{
		"I love music.",
		"We love music.",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check fFlag behaviour not failed")
	}
}

func TestSFlagBehaivourSuccess(t *testing.T) {
	data := []string{
		"A love music.",
		"I love music.",
		"C love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.sFlag = 1
	expected := []string{
		"A love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Check sFlag behaviour failed")
	}
}

func TestSFlagBehaivourFail(t *testing.T) {
	data := []string{
		"A love music.",
		"I love music.",
		"C love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	var option options
	option.sFlag = 1
	expected := []string{
		"A love music.",
		"I love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	result := checkUniqString(data, &option)
	if reflect.DeepEqual(expected, result) {
		t.Fatalf("Check sFlag behaviour not failed")
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
		cFlag: true,
		fFlag: 1,
		iFlag: true,
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
		cFlag: true,
		fFlag: 1,
		iFlag: true,
	}
	result, err := argsToOptions(args)
	if err == nil || reflect.DeepEqual(expected, result) {
		t.Fatalf("Flags check not failed")
	}
}

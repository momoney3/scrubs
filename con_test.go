package main

import "testing"

func TestConv(t *testing.T) {
	num, err := Conv("123456789")
	if err != nil {
		t.Fatal(err)
	}
	if num != 123456789 {
		t.Fatal("number dont match")
	}
}

func TestFailConv(t *testing.T) {
	_, err := conv("")
	if err == nil {
		t.Fatal(err)
	}
}

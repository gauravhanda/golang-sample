package main

import "testing"

func TestMethod(test *testing.T) {
	test.Errorf("Test Failed")
}

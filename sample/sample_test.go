package sample

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld("SaK")
	expected := "hello world, SaK"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

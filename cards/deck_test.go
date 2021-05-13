package main

import "testing"

// 't' is the test handler of go
// it will automatically called by Go testrunner
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
}

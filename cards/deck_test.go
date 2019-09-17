package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected 16,but %v", len(d))
	}

	if d[0] != "spadesoface" {
		t.Errorf("got %v", d[0])
	}

	if d[len(d)-1] != "cloveroffour" {
		t.Errorf("got %v", d[len(d)-1])
	}

}

package proverb

import "testing"

func TestSample(t *testing.T) {
	c, err := NewCollection()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	proverb := c.Sample()
	if proverb.Author == "" {
		t.Fatalf("%#v", proverb)
	}
}

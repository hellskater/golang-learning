package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("one two three\n")

	expected := 3

	actual := count(b, false, false)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, actual)
	}

}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("one two three\nline two")

	expected := 2

	actual := count(b, true, false)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, actual)
	}

}
func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("one two three\nline two")

	expected := 22

	actual := count(b, false, true)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, actual)
	}

}

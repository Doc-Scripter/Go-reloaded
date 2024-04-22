package inhouse

import (
	"testing"
)

func TestBin(t *testing.T) {
	input := "it has been 10 (bin) years "
	expected := "it has been 2 years"

	result := Bin(input)
	if result != expected {
		t.Errorf("Got %q expected %q", result, expected)
	}
}

func TestHex(t *testing.T) {
	input := "1E (hex) files were added"
	expected := "30 files were added"

	result := Hex(input)
	if result != expected {
		t.Errorf("Got  %q  expected %q", result, expected)
	}
}

func TestVowel(t *testing.T) {
	input := "There it was. A amazing rock!"
	expected := "There it was. An amazing rock!"

	result := Vowel(input)
	if result != expected {
		t.Errorf("Got %q expected %q", result, expected)
	}
}

func TestModify(t *testing.T) {
	input := "This is so exciting (up, 2)"
	expected := "This is SO EXCITING"

	result := Modify(input)
	if result != expected {
		t.Errorf("Got %q expected %q", result, expected)
	}
}

func TestPunctuation(t *testing.T) {
	input := "I am exactly how they describe me: ' awesome '"
	expected := "I am exactly how they describe me: 'awesome'"

	result := Punctuation(input)
	if result != expected {
		t.Errorf("Got Puntuation%q expected %q", result, expected)
	}
	input = "As Elton John said: ' I am the most well-known homosexual in the world '"
	expected = "As Elton John said: 'I am the most well-known homosexual in the world'"

	result = Punctuation(input)
	if result != expected {
		t.Errorf("Got %q expected %q", result, expected)
	}
}

package randomstring

import (
	"testing"
)

func TestGenerateStringFromCharset(t *testing.T) {
	tests := []struct {
		charset Charset
		length  int
		err     bool
	}{
		{Alphanumeric, 10, false},
		{Lowercase, 5, false},
		{Uppercase, 5, false},
		{Numeric, 5, false},
		{SpecialCharacters, 5, false},
		{Charset(""), 5, true},
		{Alphanumeric, -5, true},
	}

	for _, tt := range tests {
		result, err := generateStringFromCharset(tt.charset, tt.length)
		if (err != nil) != tt.err {
			t.Errorf("for charset %v and length %d, expected error: %v, got: %v", tt.charset, tt.length, tt.err, (err != nil))
		}
		if !tt.err && len(result) != tt.length {
			t.Errorf("expected length %d, got %d", tt.length, len(result))
		}
	}
}

func TestGenerateString(t *testing.T) {
	tests := []struct {
		opts GenerationOptions
		err  bool
	}{
		{GenerationOptions{Length: 10}, false},
		{GenerationOptions{Length: 10, DisableNumeric: true, DisableLowercase: true, DisableUppercase: true}, true},
		{GenerationOptions{Length: 10, DisableNumeric: true}, false},
		{GenerationOptions{Length: 10, EnableSpecialCharacter: true}, false},
		{GenerationOptions{Length: 10, CustomCharset: "abc123"}, false},
		{GenerationOptions{Length: 0}, true},
	}

	for _, tt := range tests {
		result, err := GenerateString(tt.opts)
		if (err != nil) != tt.err {
			t.Errorf("for options %+v, expected error: %v, got: %v", tt.opts, tt.err, (err != nil))
		}
		if !tt.err && len(result) != tt.opts.Length {
			t.Errorf("expected length %d, got %d", tt.opts.Length, len(result))
		}
	}
}

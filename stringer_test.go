package stringer

import (
	"testing"
)

func TestStringer_ToLower(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name string
		er   er
		want string
	}{
		{"t1", er{OriginalString: "Abc", CaseSensitive: false}, "abc"},
		{"t2", er{OriginalString: "abc", CaseSensitive: true}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.ToLower(); got.Value() != tt.want {
				t.Errorf("%s ToLower() = %v, want %v", tt.name, got.Value(), tt.want)
			}
		})
	}
}

func TestStringer_ToUpper(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name string
		er   er
		want string
	}{
		{"t1", er{OriginalString: "abc", CaseSensitive: false}, "ABC"},
		{"t2", er{OriginalString: "ABC", CaseSensitive: true}, "ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.ToUpper(); got.Value() != tt.want {
				t.Errorf("%s ToUpper() = %v, want %v", tt.name, got.Value(), tt.want)
			}
		})
	}
}

func TestStringer_HasPrefix(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name   string
		er     er
		prefix string
		want   bool
	}{
		{"t1", er{OriginalString: "Hello, World!", CaseSensitive: false}, "Hello", true},
		{"t2", er{OriginalString: " Hello, World!", CaseSensitive: false}, "Hello", false},
		{"t3", er{OriginalString: "Hello, World!", CaseSensitive: true}, "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.HasPrefix(tt.prefix); got != tt.want {
				t.Errorf("%s ToUpper() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

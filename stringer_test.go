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
		{"t2", er{OriginalString: "Abc", CaseSensitive: true}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.ToLower(); got.Value() != tt.want {
				t.Errorf("ToLower() = %v, want %v", got.Value(), tt.want)
			}
		})
	}
}

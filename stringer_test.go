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
				t.Errorf("%s ToLower() = `%v`, want `%v`", tt.name, got.Value(), tt.want)
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
				t.Errorf("%s ToUpper() = `%v`, want `%v`", tt.name, got.Value(), tt.want)
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
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: false}, "Hello", true},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, "Hello", false},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.HasPrefix(tt.prefix); got != tt.want {
				t.Errorf("%s HasPrefix() = `%v`, want `%v`", tt.name, got, tt.want)
			}
		})
	}
}

func TestStringer_HasSuffix(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name   string
		er     er
		suffix string
		want   bool
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "World!", true},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, "wOrld!", true},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, "wOrld!", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.HasSuffix(tt.suffix); got != tt.want {
				t.Errorf("%s HasSuffix() = `%v`, want `%v`", tt.name, got, tt.want)
			}
		})
	}
}

func TestStringer_TrimRight(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name   string
		er     er
		cutstr string
		want   string
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "World!", "Hello "},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, "wOrld!", " Hello "},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, "wOrld!", "Hello Wo"},
		{"t4", er{OriginalString: "Hello World!", CaseSensitive: true}, "Hello", "Hello World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.TrimRight(tt.cutstr); got.Value() != tt.want {
				t.Errorf("%s TrimRight() = `%v`, want `%v`", tt.name, got.Value(), tt.want)
			}
		})
	}
}

func TestStringer_RemoveRight(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name    string
		er      er
		custstr string
		want    string
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "World!", "Hello"},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, "wOrld!", "Hello"},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, "wOrld!", "Hello World!"},
		{"t4", er{OriginalString: "Hello World!World!", CaseSensitive: false}, "World!", "Hello"},
		{"t5", er{OriginalString: "Hello World!World!", CaseSensitive: false}, "!wOrld!", "Hello World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.RemoveRight(tt.custstr); got.Value() != tt.want {
				t.Errorf("%s RemoveRight() = `%v`, want `%v`", tt.name, got.Value(), tt.want)
			}
		})
	}
}

func TestStringer_Contains(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name    string
		er      er
		custstr string
		want    bool
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "World!", true},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, "wOrld!", true},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, "wOrld!", false},
		{"t4", er{OriginalString: "Hello World!World!", CaseSensitive: false}, "World!", true},
		{"t5", er{OriginalString: "Hello World!World!", CaseSensitive: false}, "!wOrld!", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.Contains(tt.custstr); got != tt.want {
				t.Errorf("%s Contains() = `%v`, want `%v`", tt.name, got, tt.want)
			}
		})
	}
}

func TestStringer_UpperFirst(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name string
		er   er
		want string
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "Hello World!"},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, " Hello World!"},
		{"t3", er{OriginalString: "hello World!", CaseSensitive: true}, "Hello World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.UpperFirst(); got.Value() != tt.want {
				t.Errorf("%s UpperFirst() = `%v`, want `%v`", tt.name, got.Value(), tt.want)
			}
		})
	}
}

func TestStringer_LowerFirst(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name string
		er   er
		want string
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "hello World!"},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, " Hello World!"},
		{"t3", er{OriginalString: "hello World!", CaseSensitive: true}, "hello World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.LowerFirst(); got.Value() != tt.want {
				t.Errorf("%s LowerFirst() = `%v`, want `%v`", tt.name, got.Value(), tt.want)
			}
		})
	}
}

package stringer

import (
	"strings"
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
			if got := stringer.ToLower(); got.NewValue() != tt.want {
				t.Errorf("%s ToLower() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
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
			if got := stringer.ToUpper(); got.NewValue() != tt.want {
				t.Errorf("%s ToUpper() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
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

func TestStringer_Spaceless(t *testing.T) {
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
		{"t2", er{OriginalString: " Hello     World!     ", CaseSensitive: false}, "Hello World!"},
		{"t3", er{OriginalString: "Hello    W     orld!", CaseSensitive: true}, "Hello W orld!"},
		{"t4", er{OriginalString: "Hello　　　World!", CaseSensitive: true}, "Hello World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.Spaceless(); got.NewValue() != tt.want {
				t.Errorf("%s Spaceless() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
			}
		})
	}
}

func TestStringer_TrimLeft(t *testing.T) {
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
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "Hello", " World!"},
		{"t2", er{OriginalString: "Hellohello World!", CaseSensitive: false}, "Hello!", " World!"},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, "wOrld!", "Hello World!"},
		{"t4", er{OriginalString: "hello World!", CaseSensitive: true}, "Hello", "hello World!"},
		{"t5", er{OriginalString: "111234!", CaseSensitive: true}, "123", "4!"},
		{"t6", er{OriginalString: "", CaseSensitive: true}, "123", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.TrimLeft(tt.cutstr); got.NewValue() != tt.want {
				t.Errorf("%s TrimLeft() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
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
			if got := stringer.TrimRight(tt.cutstr); got.NewValue() != tt.want {
				t.Errorf("%s TrimRight() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
			}
		})
	}
}

func TestStringer_TrimPrefix(t *testing.T) {
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
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "World!", "Hello World!"},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, "Hello", " Hello World!"},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, "hello", "Hello World!"},
		{"t4", er{OriginalString: "Hello hello World!World!", CaseSensitive: false}, "hello", " hello World!World!"},
		{"t5", er{OriginalString: "Hellohello World!World!", CaseSensitive: false}, "hello", "hello World!World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.TrimPrefix(tt.custstr); got.NewValue() != tt.want {
				t.Errorf("%s TrimPrefix() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
			}
		})
	}
}

func TestStringer_TrimSuffix(t *testing.T) {
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
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "World!", "Hello "},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, "wOrld!", " Hello "},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, "wOrld!", "Hello World!"},
		{"t4", er{OriginalString: "Hello World!World!", CaseSensitive: false}, "World!", "Hello World!"},
		{"t5", er{OriginalString: "Hello World!World!", CaseSensitive: false}, "!wOrld!", "Hello World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.TrimSuffix(tt.custstr); got.NewValue() != tt.want {
				t.Errorf("%s TrimSuffix() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
			}
		})
	}
}

func TestStringer_EqualFold(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name   string
		er     er
		string string
		want   bool
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "Hello World!", true},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, " HELLO WORLD!", true},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, " Hello World!", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.EqualFold(tt.string); got != tt.want {
				t.Errorf("%s EqualFold() = `%v`, want `%v`", tt.name, got, tt.want)
			}
		})
	}
}

func TestStringer_Equals(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name   string
		er     er
		string string
		want   bool
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "Hello World!", true},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, " HELLO WORLD!", true},
		{"t3", er{OriginalString: "Hello World!", CaseSensitive: true}, " Hello World!", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.Equals(tt.string); got != tt.want {
				t.Errorf("%s Equals() = `%v`, want `%v`", tt.name, got, tt.want)
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
			if got := stringer.UpperFirst(); got.NewValue() != tt.want {
				t.Errorf("%s UpperFirst() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
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
			if got := stringer.LowerFirst(); got.NewValue() != tt.want {
				t.Errorf("%s LowerFirst() = `%v`, want `%v`", tt.name, got.NewValue(), tt.want)
			}
		})
	}
}

func TestStringer_ContainsWord(t *testing.T) {
	type er struct {
		OriginalString string
		CaseSensitive  bool
	}
	tests := []struct {
		name string
		er   er
		word string
		want bool
	}{
		{"t1", er{OriginalString: "Hello World!", CaseSensitive: true}, "hello World!", false},
		{"t2", er{OriginalString: " Hello World!", CaseSensitive: false}, " Hello World!", true},
		{"t3", er{OriginalString: "hello World!", CaseSensitive: true}, "hello World!", true},
		{"t4", er{OriginalString: "hello World!", CaseSensitive: true}, "Hello World!", false},
		{"t5", er{OriginalString: "username", CaseSensitive: false}, "name", false},
		{"t6", er{OriginalString: "what's you name", CaseSensitive: false}, "what", false},
		{"t7", er{OriginalString: "what's you name", CaseSensitive: false}, "what's", true},
		{"t8", er{OriginalString: "0", CaseSensitive: false}, "0", true},
		{"t9", er{OriginalString: "what's you name", CaseSensitive: false}, "name", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringer := NewStringer(tt.er.OriginalString, tt.er.CaseSensitive)
			if got := stringer.ContainsWord(tt.word); got != tt.want {
				t.Errorf("%s ContainsWord() = `%v`, want `%v`", tt.name, got, tt.want)
			}
		})
	}
}

func FuzzStringer_ContainsWord(f *testing.F) {
	f.Add("hello world!", "hello")
	f.Fuzz(func(t *testing.T, originalString string, word string) {
		stringer := NewStringer(originalString, false)
		if stringer.Contains(word) != strings.Contains(strings.ToLower(originalString), strings.ToLower(word)) {
			t.Error("Not contains")
		}
	})
}

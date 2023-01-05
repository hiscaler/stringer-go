package stringer

import (
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unsafe"
)

var (
	rxSpaceless = regexp.MustCompile("\\s{2,}")
)

type Stringer struct {
	processedString      string // Processed string
	lowerProcessedString string // Processed lower string
	OriginalString       string // Original string
	CaseSensitive        bool   // Case sensitive
}

func NewStringer(s string, caseSensitive bool) *Stringer {
	return &Stringer{
		OriginalString:       s,
		CaseSensitive:        caseSensitive,
		processedString:      s,
		lowerProcessedString: strings.ToLower(s),
	}
}

func (s *Stringer) setProcessedString(str string) *Stringer {
	s.processedString = str
	s.lowerProcessedString = strings.ToLower(str)
	return s
}

func (s *Stringer) ToLower() *Stringer {
	s.setProcessedString(s.lowerProcessedString)
	return s
}

func (s *Stringer) ToUpper() *Stringer {
	s.setProcessedString(strings.ToUpper(s.processedString))
	return s
}

func (s *Stringer) HasPrefix(prefix string) bool {
	s1 := s.processedString
	if !s.CaseSensitive {
		s1 = s.lowerProcessedString
		prefix = strings.ToLower(prefix)
	}
	return strings.HasPrefix(s1, prefix)
}

func (s *Stringer) HasSuffix(suffix string) bool {
	s1 := s.processedString
	if !s.CaseSensitive {
		s1 = s.lowerProcessedString
		suffix = strings.ToLower(suffix)
	}
	return strings.HasSuffix(s1, suffix)
}

func (s *Stringer) TrimSpace() *Stringer {
	s.setProcessedString(strings.TrimSpace(s.processedString))
	return s
}

func (s *Stringer) Spaceless() *Stringer {
	str := s.processedString
	if str != "" {
		str = strings.TrimSpace(str)
	}
	if str != "" {
		str = rxSpaceless.ReplaceAllLiteralString(strings.ReplaceAll(str, "　", " "), " ")
	}
	s.setProcessedString(str)
	return s
}

func (s *Stringer) TrimLeft(cutset string) *Stringer {
	if s.processedString == "" || cutset == "" {
		return s
	}

	s1 := s.processedString
	if !s.CaseSensitive {
		s1 = s.lowerProcessedString
		cutset = strings.ToLower(cutset)
	}
	s1 = strings.TrimLeft(s1, cutset)
	s.setProcessedString(s.processedString[len(s.processedString)-len(s1):])
	return s
}

func (s *Stringer) TrimRight(cutset string) *Stringer {
	if s.processedString == "" || cutset == "" {
		return s
	}

	s1 := s.processedString
	if !s.CaseSensitive {
		s1 = s.lowerProcessedString
		cutset = strings.ToLower(cutset)
	}
	s1 = strings.TrimRight(s1, cutset)
	s.setProcessedString(s.processedString[0:len(s1)])
	return s
}

func (s *Stringer) Index(substr string) int {
	s1 := s.processedString
	if !s.CaseSensitive {
		s1 = s.lowerProcessedString
		substr = strings.ToLower(substr)
	}
	return strings.Index(s1, substr)
}

func (s *Stringer) LastIndex(substr string) int {
	s1 := s.processedString
	if !s.CaseSensitive {
		s1 = s.lowerProcessedString
		substr = strings.ToLower(substr)
	}
	return strings.LastIndex(s1, substr)
}

func (s *Stringer) TrimPrefix(prefix string) *Stringer {
	s1 := s.processedString
	if !s.CaseSensitive {
		prefix = strings.ToLower(prefix)
		s1 = s.lowerProcessedString
	}
	n := len(s1)
	s1 = strings.TrimPrefix(s1, prefix)
	if s1 == "" {
		s.setProcessedString("")
	} else {
		s.setProcessedString(s.processedString[n-len(s1):])
	}

	return s
}

func (s *Stringer) TrimSuffix(suffix string) *Stringer {
	s1 := s.processedString
	if !s.CaseSensitive {
		suffix = strings.ToLower(suffix)
		s1 = s.lowerProcessedString
	}
	s.setProcessedString(s.processedString[0:len(strings.TrimSuffix(s1, suffix))])
	return s
}

func (s *Stringer) EqualFold(t string) bool {
	return s.lowerProcessedString == strings.ToLower(t)
}

func (s *Stringer) IsEmpty() bool {
	return len(s.processedString) == 0
}

func (s *Stringer) IsBlank() bool {
	return len(s.processedString) == 0 || strings.TrimSpace(s.processedString) == ""
}

func (s *Stringer) Contains(substr string) bool {
	return s.Index(substr) >= 0
}

func (s *Stringer) ContainsWord(word string) bool {
	expr := `(^|([\s\t\n]+))(` + word + `)($|([\s\t\n]+))`
	if !s.CaseSensitive {
		expr = "(?i)" + expr
	}
	return regexp.MustCompile(expr).MatchString(s.processedString)
}

func (s *Stringer) UpperFirst() *Stringer {
	r := []rune(s.processedString)
	if len(s.processedString) > 0 && unicode.IsLetter(r[0]) && unicode.IsLower(r[0]) {
		r[0] -= 32
		s.setProcessedString(string(r))
	}
	return s
}

func (s *Stringer) LowerFirst() *Stringer {
	r := []rune(s.processedString)
	if len(s.processedString) > 0 && unicode.IsLetter(r[0]) && unicode.IsUpper(r[0]) {
		r[0] += 32
		s.setProcessedString(string(r))
	}
	return s
}

func (s *Stringer) Value() string {
	return s.processedString
}

func (s *Stringer) ToByte() []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s.processedString))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

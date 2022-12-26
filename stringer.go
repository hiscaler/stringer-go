package stringer

import (
	"reflect"
	"strings"
	"unsafe"
)

type Stringer struct {
	processedString      string // Processed string
	lowerProcessedString string // Processed lower string
	OriginalString       string // Original string
	CaseSensitive        bool   // Case Sensitive
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

func (s *Stringer) TrimRight(cutset string) *Stringer {
	s1 := s.processedString
	if !s.CaseSensitive {
		s1 = s.lowerProcessedString
	}
	s.setProcessedString(strings.TrimRight(s1, cutset))
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

func (s *Stringer) RemoveRight(str string) *Stringer {
	for {
		if !s.HasSuffix(str) {
			return s
		}
		index := s.LastIndex(str)
		if index == -1 {
			return s
		}
		s.setProcessedString(strings.TrimSpace(s.processedString[0:index]))
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

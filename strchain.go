package strutil

import (
	"strings"

	"github.com/duke-git/lancet/strutil"
)

type strchain struct {
	string
}

func NewStrChain(s string) *strchain {
	sc := new(strchain)
	sc.string = s
	return sc
}

func (s *strchain) TrimSpace() *strchain {
	s.string = strings.TrimSpace(s.string)
	return s
}

func (s *strchain) ToUpper() *strchain {
	s.string = strings.ToUpper(s.string)
	return s
}

func (s *strchain) ToLower() *strchain {
	s.string = strings.ToLower(s.string)
	return s
}

func (s *strchain) ReplaceAll(old, new string) *strchain {
	s.string = strings.ReplaceAll(s.string, old, new)
	return s
}

// 返回源字符串中特定字符串首次出现时的位置之后的子字符串。
func (s *strchain) After(char string) *strchain {
	s.string = strutil.After(s.string, char)
	return s
}

// 返回源字符串中指定字符串最后一次出现时的位置之后的子字符串。
func (s *strchain) AfterLast(char string) *strchain {
	s.string = strutil.AfterLast(s.string, char)
	return s
}

// 返回源字符串中指定字符串第一次出现时的位置之前的子字符串。
func (s *strchain) Before(char string) *strchain {
	s.string = strutil.Before(s.string, char)
	return s
}

// 返回源字符串中指定字符串最后一次出现时的位置之前的子字符串。
func (s *strchain) BeforeLast(char string) *strchain {
	s.string = strutil.BeforeLast(s.string, char)
	return s
}

func (s *strchain) Containers(char string) bool {
	return strings.Contains(s.string, char)
}

func (s *strchain) IsBlank() bool {
	return strutil.IsBlank(s.string)
}

func (s *strchain) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.string, prefix)
}

func (s *strchain) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.string, suffix)
}

func (s *strchain) String() string {
	return s.string
}

func (s *strchain) Bytes() []byte {
	return []byte(s.string)
}

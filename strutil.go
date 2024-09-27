package strutil

import (
	"strings"

	"github.com/duke-git/lancet/strutil"
)

type StrUtil struct {
	string
}

func NewStrUtil(s string) *StrUtil {
	su := new(StrUtil)
	su.string = s
	return su
}

func (s *StrUtil) TrimSpace() *StrUtil {
	s.string = strings.TrimSpace(s.string)
	return s
}

func (s *StrUtil) ToUpper() *StrUtil {
	s.string = strings.ToUpper(s.string)
	return s
}

func (s *StrUtil) ToLower() *StrUtil {
	s.string = strings.ToLower(s.string)
	return s
}

func (s *StrUtil) ReplaceAll(old, new string) *StrUtil {
	s.string = strings.ReplaceAll(s.string, old, new)
	return s
}

// 返回源字符串中特定字符串首次出现时的位置之后的子字符串。
func (s *StrUtil) After(char string) *StrUtil {
	s.string = strutil.After(s.string, char)
	return s
}

// 返回源字符串中指定字符串最后一次出现时的位置之后的子字符串。
func (s *StrUtil) AfterLast(char string) *StrUtil {
	s.string = strutil.AfterLast(s.string, char)
	return s
}

// 返回源字符串中指定字符串第一次出现时的位置之前的子字符串。
func (s *StrUtil) Before(char string) *StrUtil {
	s.string = strutil.Before(s.string, char)
	return s
}

// 返回源字符串中指定字符串最后一次出现时的位置之前的子字符串。
func (s *StrUtil) BeforeLast(char string) *StrUtil {
	s.string = strutil.BeforeLast(s.string, char)
	return s
}

func (s *StrUtil) Containers(char string) bool {
	return strings.Contains(s.string, char)
}

func (s *StrUtil) IsBlank() bool {
	return strutil.IsBlank(s.string)
}

func (s *StrUtil) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.string, prefix)
}

func (s *StrUtil) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.string, suffix)
}

func (s *StrUtil) String() string {
	return s.string
}

func (s *StrUtil) Bytes() []byte {
	return []byte(s.string)
}

package word

import (
	"strings"
	"unicode"
)

func GetRune(str string, index int32) string {
	return string([]rune(str)[index])
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func Underscore_To_UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func Underscore_To_LowerCamelCase(s string) string {
	s = Underscore_To_UpperCamelCase(s)
	return strings.ToLower(GetRune(s, 0)) + s[1:]
}

func CamelCase_To_Underscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

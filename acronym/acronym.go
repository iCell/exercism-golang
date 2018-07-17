package acronym

import "strings"

func Abbreviate(s string) string {
	var result []rune
	var splits = strings.FieldsFunc(s, func (r rune) bool {
		return string(r) == " " || string(r) == "-"
	})
	for _, v := range splits {
		if len(v) > 0 {
			str := []rune(strings.ToUpper(v))
			result = append(result, str[0])
		}
	}
	return string(result)
}

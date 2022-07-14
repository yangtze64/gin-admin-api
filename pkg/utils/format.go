package utils

import "strings"

var (
	TimeFormatMap = map[string]string{
		"Y": "2006",
		"y": "06",
		"m": "01",
		"n": "1",
		"d": "02",
		"j": "2",
		"H": "15",
		"h": "03",
		"g": "3",
		"i": "04",
		"s": "05",
	}
)

func GetTimeFormatStr(formatStr string) string {
	for k, v := range TimeFormatMap {
		formatStr = strings.Replace(formatStr, k, v, -1)
	}
	return formatStr
}

package utils

import "testing"

func TestGetTimeFormatStr(t *testing.T) {
	t.Log(GetTimeFormatStr("Y-m-d H:i:s"))
	t.Log(GetTimeFormatStr("Ymd His"))
	t.Log(GetTimeFormatStr("Y-n-j"))
}

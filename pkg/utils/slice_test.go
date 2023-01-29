package utils

import (
	"fmt"
	"testing"
)

func TestSliceMap(t *testing.T) {
	slice := []interface{}{"aa", "bb", "cc"}
	slicenew := SliceMap(slice, func(i interface{}) interface{} {
		return i.(string) + "1"
	})
	fmt.Printf("%#+v\n", slicenew)
}

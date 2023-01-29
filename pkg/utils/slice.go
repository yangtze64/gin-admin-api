package utils

func SliceMap(slice []interface{}, cb func(interface{}) interface{}) []interface{} {
	res := make([]interface{}, len(slice))
	copy(res, slice)
	for i, v := range slice {
		res[i] = cb(v)
	}
	return res
}

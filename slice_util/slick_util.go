package slice_util

// 把切片复制一份
func Clone(slice []interface{}) []interface{} {
	if slice == nil {
		return nil
	}
	newSlice := make([]interface{}, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// 对切片去重
func Distinct(slice []interface{}) []interface{} {
	itemMap := make(map[interface{}]struct{})
	newSlice := make([]interface{}, 0)
	for _, item := range slice {
		if _, exists := itemMap[item]; exists {
			continue
		}
		newSlice = append(newSlice, item)
		itemMap[item] = struct{}{} //0字节
	}
	return newSlice
}

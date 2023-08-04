package common

import "reflect"

// 判断一个interface是否为空
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// 获取interface的string值
func GetStringVal(i interface{}) string {
	if !IsNil(i) {
		return reflect.ValueOf(i).String()
	}
	return ""
}

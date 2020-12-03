package common

import "reflect"

func IsStructEmpty(i interface{}) bool {
	return reflect.ValueOf(i).IsZero()
}

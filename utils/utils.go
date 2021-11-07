// utils.go kee > 2021/11/03

package utils

import (
	"reflect"
)

func GetTypeOf(options []interface{}, name string) (int, interface{}) {
	for i, opt := range options {
		switch reflect.TypeOf(opt).Name() {
		case name:
			return i, opt
		}
	}
	return 0, nil
}

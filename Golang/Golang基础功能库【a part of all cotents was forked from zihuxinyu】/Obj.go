package Library

import "reflect"

//检查是否为指针
func isPtr(obj interface {}) (bool){
	return reflect.TypeOf(obj).Kind()==reflect.Ptr
}

package main

import (
	"fmt"
	"reflect"
)

// AppendMember exported
func AppendMember(list interface{}, member interface{}) interface{} {
	found := false
	lt, mt := reflect.TypeOf(list), reflect.TypeOf(member)
	if lt.Kind() != reflect.Slice {
		fmt.Printf("list type not Slice")
	}

	if lt.Elem() != mt {
		fmt.Printf("list type not same as member type")
	}

	listValue := reflect.ValueOf(list)
	for i := 0; i < listValue.Len(); i++ {
		if listValue.Index(i).Interface() == member {
			found = true
			break
		}
	}
	if !found {
		reflect.ValueOf(member).Interface()
		newList := reflect.Append(reflect.ValueOf(list), reflect.ValueOf(member))
		return newList.Interface()
	}
	return list
}

func main() {
	x := []int{5, 10, 100}
	x = AppendMember(x, 100).([]int)
	fmt.Println(x)
}

package main

import (
	"reflect"
	"testing"
)

func TestAppendMember(t *testing.T) {
	x := []int{5, 10, 100}
	exp := []int{5, 10, 100, 1000}
	x = AppendMember(x, 1000).([]int)
	if !reflect.DeepEqual(x, exp) {
		t.Errorf("Not expected, recvd %v exp %v", x, exp)
	}
}

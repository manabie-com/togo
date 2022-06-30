package utils

import (
	"reflect"
	"testing"
)

func TestStr2Uint32(t *testing.T) {
	var predict uint32 = 100
	actual, err := Str2Uint32("100")

	if err != nil {
		t.Errorf("expected %v, but got %v", nil, err)
	}
	if predict != actual {
		t.Errorf("expected %v, but got %v", predict, actual)
	}
	if reflect.TypeOf(actual) != reflect.TypeOf(predict) {
		t.Errorf("expected %v, but got %v", "uint32", reflect.TypeOf(actual))
	}
}

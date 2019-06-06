package arethereduplicate

import (
	"fmt"
	"reflect"
)

func areThereDuplicates(vals interface{}) (bool, error) {
	s := reflect.ValueOf(vals)

	unique := make(map[interface{}]int)

	switch reflect.TypeOf(vals).Elem().Kind() {
	case reflect.Int:
		for i := 0; i < s.Len(); i++ {
			a := s.Index(i).Int()
			if unique[a] >= 1 {
				return true, nil
			}
			unique[a]++
		}
	case reflect.String:
		for i := 0; i < s.Len(); i++ {
			a := s.Index(i).String()
			if unique[a] >= 1 {
				return true, nil
			}
			unique[a]++
		}
	default:
		return false, fmt.Errorf("Type not supported.")
	}

	return false, nil
}

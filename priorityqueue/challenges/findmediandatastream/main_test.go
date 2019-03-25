package findmedian

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddNum(t *testing.T) {
	tests := []struct {
		values []int
	}{
		{
			[]int{2, 3, 4},
		},
	}

	for _, tt := range tests {
		mf := Constructor()

		for _, n := range tt.values {
			mf.AddNum(n)
		}

		got := mf.elements

		if !reflect.DeepEqual(got, tt.values) {
			fmt.Errorf("Want %v, got %v", tt.values, got)
		}
	}
}

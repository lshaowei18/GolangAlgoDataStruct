package hashtable

import (
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		key  string
		want int
	}{
		{"hello", 40},
		{"goodbye", 23},
		{"cyan", 23},
		{"hi", 45},
	}

	for _, tt := range tests {
		ht := HashTable{}
		got := ht.Hash(tt.key)

		if got != tt.want {
			t.Errorf("Got %v, want %v", got, tt.want)
		}
	}
}

func TestSet(t *testing.T) {
	t.Run("Different keys", func(t *testing.T) {
		values := []string{"hello", "goodbye", "hi"}

		ht := HashTable{}

		got := []string{}
		for _, v := range values {
			k := ht.Hash(v)
			ht.Set(k, v)
			got = append(got, ht.keyMap[k][0])
		}

		if !reflect.DeepEqual(got, values) {
			t.Errorf("Got %v, want %v", got, values)
		}
	})

	t.Run("Same keys", func(t *testing.T) {
		values := []string{"goodbye", "cyan"}

		ht := HashTable{}

		var k int

		for _, v := range values {
			k = ht.Hash(v)
			ht.Set(k, v)
		}

		got := ht.keyMap[k]

		if !reflect.DeepEqual(got, values) {
			t.Errorf("Got %v, want %v", got, values)
		}
	})
}

func TestGet(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		key    string
		want   string
	}{
		{
			"Different hash keys",
			[]string{"hello", "goodbye", "hi"},
			"hello",
			"hello",
		},
		{
			"Invalid key",
			[]string{"hello", "goodbye", "hi"},
			"chicken",
			"",
		},
		{
			"Same hash keys",
			[]string{"cyan", "goodbye"},
			"cyan",
			"cyan",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := HashTable{}

			for _, v := range tt.values {
				ht.Set(ht.Hash(v), v)
			}

			got := ht.Get(tt.key)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

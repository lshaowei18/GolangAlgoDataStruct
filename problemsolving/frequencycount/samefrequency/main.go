package samefrequency

import (
	"reflect"
	"strconv"
)

func sameFrequency(n1, n2 int) bool {
	//Make two maps
	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	//Convert ints to strings
	s1 := strconv.Itoa(n1)
	s2 := strconv.Itoa(n2)

	//Loop through and add char to map, with its value as the count
	for _, c := range s1 {
		count1[c]++
	}

	for _, c := range s2 {
		count2[c]++
	}

	//Return true if maps are equal, false if it is not
	return reflect.DeepEqual(count1, count2)
}

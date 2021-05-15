package compare

import (
	"reflect"
	"strings"
)

// DiffUnorderedSlice compares two arrays or slices, ignoring their order.
// Groups and the elements that are in both, only in A and only in B slices.
// Duplicate elements are treated individually, meaning the number of occurrences is verified.
// Elements are compared using the provided equal function.
func DiffUnorderedSlice(listA, listB interface{}, compareFunc func(a, b interface{}) bool) (common []interface{}, onlyA []interface{}, onlyB []interface{}) {
	// based on testify.assert (ElementsMatch)

	if !isSliceOrArray(listA) || !isSliceOrArray(listB) {
		panic("invalid arguments: slices or arrays expected")
	}

	valA := reflect.ValueOf(listA)
	valB := reflect.ValueOf(listB)
	lenA := valA.Len()
	lenB := valB.Len()

	visited := make([]bool, lenB) // for each element in listB if it has already been found in listA

	for i := 0; i < lenA; i++ {
		elemA := valA.Index(i).Interface()

		found := false
		for j := 0; j < lenB; j++ {
			if visited[j] {
				continue
			}
			elemB := valB.Index(j).Interface()
			if compareFunc(elemA, elemB) {
				common = append(common, elemA)
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			onlyA = append(onlyA, elemA)
		}
	}
	for j := 0; j < lenB; j++ {
		if visited[j] {
			continue
		}
		elemB := valB.Index(j).Interface()
		onlyB = append(onlyB, elemB)
	}
	return
}

func isSliceOrArray(list interface{}) bool {
	kind := reflect.TypeOf(list).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

// String is a compare function for two strings
func String(a, b interface{}) bool {
	return a.(string) == b.(string)
}

// SubString is a compare function that reports whether 'a' is a substring of 'b'.
func SubString(a, b interface{}) bool {
	return strings.Contains(b.(string), a.(string))
}

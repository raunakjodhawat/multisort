// Package multisort extends the official Go Package to expand the functionality provided by Sort Package
package multisort

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

// Less, Len and Swap function are required by the Sort interface
// Less function, checks if value at ith index is less than jth index
func (ms multiSortSlice) Less(i, j int) bool {
	return getLessValue(reflect.ValueOf(ms[i]).FieldByName(getKey()), reflect.ValueOf(ms[j]).FieldByName(getKey()))
}

// Len returns the length of the slice (interface)
func (ms multiSortSlice) Len() int {
	return len(ms)
}

// Swap swaps the element at index i, j
func (ms multiSortSlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

// multiSortInterface, holds the incoming interface
type multiSortInterface interface{}
// multiSortSlice, is used to []interface{} to []multiSortInterface{}
type multiSortSlice []multiSortInterface

// MultiSorted takes in a unsortedSlice interface and the inputSortKeys []string, ascendingSortOrder []bool
// returns the sorted slice based on the keys specified by sortKeys and ascendingSortOrder values
// returns []multiSortInterface, and error
// []multiSortInterface can be taken as interface in Clients program and can be converted to desired type
func MultiSorted(unsortedSlice interface{}, inputSortKeys []string, ascendingSortOrder []bool) ([]multiSortInterface, error) {
	// return if not a slice
	if reflect.TypeOf(unsortedSlice).Kind() != reflect.Slice {
		return nil, fmt.Errorf("input is not a slice")
	}

	// Anonymous function to launch One Go routines, for concurrent processing of non related processing
	var wg sync.WaitGroup
	wg.Add(1)
	// Launch a Go routine, for parallel execution
	go copyKeys(&wg, inputSortKeys)
	// Stores the unsorted slice
	var ms multiSortSlice
	// Copy the input to the multiSortSlice T
	ms = copyUnsortedSliceToMultiSorted(unsortedSlice)
	// By default sort by the order in which Keys is received.
	// Then By the order (if present) in ascendingSortOrder, else (Ascending order as default ordering)
	wg.Wait()

	// For every key in the sortKeys input
	for i := range sortKeys {
		// Change the index of current Key. sortKeys are made available globally, through the copyKeys routine
		currentKeyIndex = i
		// Check if there's a value with the current global key
		reflectValue := reflect.ValueOf(ms[0]).FieldByName(getKey())
		if !reflectValue.IsValid() {
			return ms, fmt.Errorf("%v, not present (as key) on the input slice", getKey())
		}
		// sort on the basis of default ordering if no user defined sort order is present
		if i < len(ascendingSortOrder) && ascendingSortOrder[i] {
			sort.Sort(ms)
		} else {
			sort.Sort(sort.Reverse(ms))
		}
	}
	// return the sorted slice as []multiSortInterface, which can be converted to user defined Type
	return ms, nil
}

// currentKeyIndex stored the current index, as part of global context
var currentKeyIndex int
// inputSortKeys are copied to global variables. Because, the sort Interface does not implement a method that takes in the key
var sortKeys []string

// copyKeys, copies the input keys to a global object
func copyKeys(wg *sync.WaitGroup, inputSortKeys []string) {
	defer wg.Done()
	// Copy SortKeys to global object
	for _, key := range inputSortKeys {
		sortKeys = append(sortKeys, key)
	}
}

// copyUnsortedSliceToMultiSorted, converts T.interface{} -> T.multiSortSlice (T.[]multiSortInterface)
func copyUnsortedSliceToMultiSorted(unsortedSlice interface{}) multiSortSlice {
	var sortSlice multiSortSlice
	reflectCopy := reflect.Indirect(reflect.ValueOf(unsortedSlice))
	for i := 0; i < reflectCopy.Len(); i++ {
		sortSlice = append(sortSlice, reflectCopy.Index(i).Interface())
	}
	return sortSlice
}

// getLessValue, based on the type of values, for which we are sorting. This function converts it into equivalent types
// It performs the calculation required by Less function in sort interface and returns a Bool
func getLessValue(msI, msJ reflect.Value) bool {
	switch msI.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return msI.Int() < msJ.Int()
	case reflect.Float32, reflect.Float64:
		return msI.Float() < msJ.Float()
	default:
		return msI.String() < msJ.String()
	}
}

// getKey, returns the current key by which sorting is taking place
func getKey() string {
	return sortKeys[currentKeyIndex]
}

// Help, provides a use case of how to convert the T.([]interface) into the required type
// desiredType is the type which the client sends in the data
// The motive being, that developer does not have to go back to github to figure out what to do with the sorted slice
func Help() string {
	return `outputSlice, err := MultiSorted(inputSlice, inputKeys, inputOrder)
	for i := range outputSlice {
		outputSlice[i] = outputSlice[i].(desiredType)
	}`
}
// Package multisort extends the official Go Package to expand the functionality provided by Sort Package
package multisort

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

func (ms multiSortSlice) Less(i, j int) bool{
	return getLessValue(reflect.ValueOf(ms[i]).FieldByName(getKey()), reflect.ValueOf(ms[j]).FieldByName(getKey()))
}

// Len returns the length of the slice (interface)
func (ms multiSortSlice) Len() int{
	return len(ms)
}

// Swap swaps the element at index i, j
func (ms multiSortSlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

type multiSortInterface interface{}
type multiSortSlice []multiSortInterface

// MultiSort takes in a data *interface and the dataSlice *interface
// it returns the sorted slice based on the keys specified by sortKeys and ascendingSortOrder values
// returns an error, if not nil
func MultiSort(unsortedSlice interface{}, inputSortKeys []string, ascendingSortOrder []bool) ([]multiSortInterface, error){
	// return if not a slice
	if reflect.TypeOf(unsortedSlice).Kind() != reflect.Slice{
		return nil, fmt.Errorf("input is not a slice")
	}

	var ms multiSortSlice
	// Anonymous function to launch One Go routines, for concurrent processing of non related processing
	var wg sync.WaitGroup
	wg.Add(1)
	go copyKeys(&wg, inputSortKeys)
	wg.Wait()

	ms = copyUnsortedSliceToMultiSort(unsortedSlice)
	// By default sort by the order in which Keys is received. Then By the order (if present) in ascendingSortOrder, else (Ascending order as default ordering)
	// Iterate on the sortKeys
	for i, _ := range sortKeys {
		currentKeyIndex = i
		reflectValue := reflect.ValueOf(ms[0]).FieldByName(getKey())
		if !reflectValue.IsValid() {
			return ms, fmt.Errorf("%v, not present (as key) on the input slice", getKey())
		}
		if i < len(ascendingSortOrder) && ascendingSortOrder[i] {
			sort.Sort(ms)
		} else {
			sort.Sort(sort.Reverse(ms))
		}
	}
	return ms, nil
}

var currentKeyIndex int
var sortKeys []string

func copyKeys(wg *sync.WaitGroup, inputSortKeys []string){
	defer wg.Done()
	// Copy SortKeys to global object
	for _, key := range inputSortKeys {
		sortKeys = append(sortKeys, key)
	}
}

func copyUnsortedSliceToMultiSort(unsortedSlice interface{}) multiSortSlice{
	var sortSlice multiSortSlice
	reflectCopy := reflect.Indirect(reflect.ValueOf(unsortedSlice))
	for i:= 0; i< reflectCopy.Len() ;i++ {
		sortSlice = append(sortSlice, reflectCopy.Index(i).Interface())
	}
	return sortSlice
}

func getLessValue(msI, msJ reflect.Value) bool{
	switch msI.Kind(){
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,reflect.Uint64:
		return msI.Int() < msJ.Int()
	case reflect.Float32, reflect.Float64:
		return msI.Float() < msJ.Float()
	default:
		return msI.String() < msJ.String()
	}
}

func getKey() string {
	return sortKeys[currentKeyIndex]
}
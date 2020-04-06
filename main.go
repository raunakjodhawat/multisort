// Package multisort extends the official Go Package to expand the functionality provided by Sort Package
package multisort

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

func (ms multiSortSlice) Less(i, j int) bool{
	reflectValue := reflect.ValueOf(ms[0]).FieldByName(GetKey())
	if !reflectValue.IsValid() {
		return true
	}
	return GetLessValue(reflect.ValueOf(ms[i]).FieldByName(GetKey()), reflect.ValueOf(ms[j]).FieldByName(GetKey()))
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
func MultiSort(unsortedSlice interface{}, sortKeys []string, ascendingSortOrder []bool){
	// return if not a slice
	if reflect.TypeOf(unsortedSlice).Kind() != reflect.Slice{
		fmt.Errorf("input is not a slice")
	}
	// TODO: Application does not throw any fatal error

	var ms multiSortSlice
	// Anonymous function to launch three Go routines, for concurrent processing of non related processing
	var wg sync.WaitGroup
	wg.Add(2)
	go normalizeAscendingSortOrderSlice(&wg, sortKeys, ascendingSortOrder)
	go copyKeys(&wg, sortKeys)
	wg.Wait()

	ms = copyUnsortedSliceToMultiSort(unsortedSlice)
	fmt.Println(ms)
	// By default sort by the order in which Keys is received. Then By the order (if present) in ascendingSortOrder, else (Ascending order as default ordering)
	// Iterate on the sortKeys
	for i, _ := range SortKeys {
		CurrentKeyIndex = i
		sort.Sort(ms)
		// ms
		fmt.Println(ms)
		// TODO: send correct slice
	}
	//
}

func normalizeAscendingSortOrderSlice(wg *sync.WaitGroup, sortKeys []string, ascendingSortOrder []bool){
	defer wg.Done()
	// check len of sortKeys and sortOrder
	// If length of sortKeys < ascendingSortOrder. Anything after the length of sortKeys, for ascendingSortOrder is ignored.
	// If length of sortKeys > ascendingSortOrder. True is appended to ascendingSortOrder, until its length become equal to sortKeys
	// Default sort direction is Ascending
	if len(sortKeys) != len(ascendingSortOrder) {
		for len(sortKeys) >= len(ascendingSortOrder) {
			ascendingSortOrder = append(ascendingSortOrder, true)
		}
	}
}

func copyKeys(wg *sync.WaitGroup, sortKeys []string){
	defer wg.Done()
	// Copy SortKeys to global object
	for _, key := range sortKeys {
		SortKeys = append(SortKeys, key)
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

func GetLessValue(msI, msJ reflect.Value) bool{
	switch msI.Kind(){
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,reflect.Uint64:
		return msI.Int() < msJ.Int()
	case reflect.Float32, reflect.Float64:
		return msI.Float() < msJ.Float()
	default:
		return msI.String() < msJ.String()
	}
}

var CurrentKeyIndex int
var SortKeys []string

func GetKey() string {
	return SortKeys[CurrentKeyIndex]
}
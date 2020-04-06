// Package multisort extends the official Go Package to expand the functionality provided by Sort Package
//package multisort
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

// Type multiSortInterface implements Sort interface
type multiSortStruct struct{
	Name string
	Age int
}

type multiSortSlice []multiSortStruct
// Less perform swap operation
func (ms multiSortSlice) CustomLess(i, j int, key string) bool{
	reflectValue := reflect.ValueOf(ms[0]).FieldByName(key)
	if !reflectValue.IsValid() || reflectValue.IsNil() {
		return true
	}
	return reflect.ValueOf(ms[i]).FieldByName(key).String() > reflect.ValueOf(ms[j]).FieldByName(key).String()
}

func getKey() string {
	return "Name"
}

func (ms multiSortSlice) Less(i, j int) bool{
	//reflectValue := reflect.ValueOf(ms[0]).FieldByName(getKey())
	//if !reflectValue.IsValid() || reflectValue.IsNil() {
	//	return true
	//}
	return reflect.ValueOf(ms[i]).FieldByName(getKey()).String() < reflect.ValueOf(ms[j]).FieldByName(getKey()).String()
}

// Len returns the length of the slice (interface)
func (ms multiSortSlice) Len() int{
	return len(ms)
}

// Swap swaps the element at index i, j
func (ms multiSortSlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

// Main function, for testing the functionality by the developer
func main(){
	ms1 := multiSortStruct{
		Name: "Joe",
		Age: 26,
	}
	ms2 := multiSortStruct{
		Name: "Avin",
		Age: 35,
	}

	ms := []multiSortStruct{ms1, ms2}
	fmt.Println(ms)
	MulTest(ms, []string{"Name"}, []bool{true})
	fmt.Println(ms)
	//MultiSort(personSlice, []string{"name"}, []bool{true})
}

func MulTest(ms multiSortSlice, sortKeys []string, ascendingSortOrder []bool){
	sort.Sort(ms)
}
// MultiSort takes in a data *interface and the dataSlice *interface
// it returns the sorted slice based on the keys specified by sortKeys and ascendingSortOrder values
// returns an error, if not nil
func MultiSort(data interface{}, sortKeys []string, ascendingSortOrder []bool) error{
	// return if not a slice
	if reflect.TypeOf(data).Kind() != reflect.Slice{
		return fmt.Errorf("input is not a slice")
	}
	// Convert interface to slice
	reflectCopy := reflect.ValueOf(data)
	reflectSlice:= make([]interface{}, reflectCopy.Len())
	for i, _ := range reflectSlice {
		reflectSlice[i] = reflectCopy.Index(i).Interface()
	}
	t := reflect.TypeOf(data).Elem()
	ms := reflect.New(t).Elem().Interface()
	fmt.Println(reflect.TypeOf(ms))
	// check len of sortKeys and sortOrder
		// If length of sortKeys < ascendingSortOrder. Anything after the length of sortKeys, for ascendingSortOrder is ignored.
		// If length of sortKeys > ascendingSortOrder. True is appended to ascendingSortOrder, until its length become equal to sortKeys
		// Default sort direction is Ascending
	if len(sortKeys) != len(ascendingSortOrder) {
		for len(sortKeys) >= len(ascendingSortOrder) {
			ascendingSortOrder = append(ascendingSortOrder, true)
		}
	}
	// By default sort by the order in which Keys is received. Then By the order (if present) in ascendingSortOrder, else (Ascending order as default ordering)
	// Iterate on the sortKeys
	for _, key := range sortKeys {
		typeOfObject := reflect.TypeOf(key)

		switch typeOfObject.Kind() {
		case reflect.String:
			if reflect.TypeOf(reflectSlice[0]).Kind() == reflect.Struct {
				//sort.Sort(data.(multiSortInterface))
			}
		case reflect.Slice:
			//
		}
	}
	return nil
}

func createStructType(data []byte, dataStruct interface{}) interface{} {
	json.Unmarshal(data, &dataStruct)
	reflectCopy := reflect.ValueOf(data)
	reflectSlice:= make([]interface{}, reflectCopy.Len())
	for i, _ := range reflectSlice {
		reflectSlice[i] = reflectCopy.Index(i).Interface()
	}
	type dataType struct{}
	for _, v := range reflectSlice {
		fmt.Println(reflect.Indirect(reflect.ValueOf(v)).Field(0).Type())
	}
	fmt.Println(reflect.Indirect(reflect.ValueOf(data[0])).Field(0).Type().Name())
	return dataType{}
}
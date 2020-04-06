// Package multisort extends the official Go Package to expand the functionality provided by Sort Package
//package multisort
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Type multiSortInterface implements Sort interface
type multiSortInterface []interface{}

// Less perform swap operation
func (ms multiSortInterface) Less(i, j int) bool{
	return true
	//return ms[i] > ms[j]
}

// Len returns the length of the slice (interface)
func (ms multiSortInterface) Len() int{
	return len(ms)
}

// Swap swaps the element at index i, j
func (ms multiSortInterface) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

// Main function, for testing the functionality by the developer
func main(){
	type person struct{
		name string
		age int
	}
	p1 := person{
		name: "Joe",
		age: 26,
	}
	p2 := person{
		name: "Avin",
		age: 35,
	}


	personSlice := []person{p1, p2}

	createStructType(json.Marshal(personSlice), person{})
	MultiSort(personSlice, []string{"name"}, []bool{true})
	fmt.Println(personSlice)
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
// MultiSort takes in a data *interface and the dataSlice *interface
// it returns the sorted slice based on the keys specified by sortKeys and ascendingSortOrder values
// returns an error, if not nil
func MultiSort(data interface{}, sortKeys []string, ascendingSortOrder []bool) error{
	dataStruct := createStructType(data)
	fmt.Println(dataStruct, "leave")
	// return if not a slice
	if reflect.TypeOf(data).Kind() != reflect.Slice{
		return fmt.Errorf("input is not a slice")
	}
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

		// Convert interface to slice
		reflectCopy := reflect.ValueOf(data)
		reflectSlice:= make([]interface{}, reflectCopy.Len())
		for i, _ := range reflectSlice {
			reflectSlice[i] = reflectCopy.Index(i).Interface()
		}

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
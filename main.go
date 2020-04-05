// Package multisort extends the official Go Package to expand the functionality provided by Sort Package
//package multisort
package main

import "fmt"

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
	fmt.Println(personSlice)
	//MultiSort(&person, &personSlice, []string{"name"}, []bool{true})
}



// MultiSort takes in a data interface and the dataSlice object, it returns the sorted slice based on the keys specified by sortKeys and ascendingSortOrder values
func MultiSort(data *struct{}, dataSlice *[]interface{} , sortKeys []string, ascendingSortOrder []bool){
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
		fmt.Println("key", key)
	}
}
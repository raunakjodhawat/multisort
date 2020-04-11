// Package MultiSort provides gives the ability to sort a slice on the basis of multiple keys
package multisort

import (
	"fmt"
	"testing"
)

// multiSortExamplePerson type, is used as an example type. It has two field, Name and Age
type multiSortExamplePerson struct {
	Name string
	Age  int
}

// makeInput Function is a utility function, used by the Examples function to remove code redundancy
// it creates a slice of type multiSortExamplePerson, again multiSortExamplePerson here is used as a Generic type T
func makeInput() []multiSortExamplePerson {
	p1 := multiSortExamplePerson{
		Name: "Joe",
		Age:  26,
	}
	p2 := multiSortExamplePerson{
		Name: "Azin",
		Age:  14,
	}
	p3 := multiSortExamplePerson{
		Name: "Bold",
		Age:  11,
	}
	p4 := multiSortExamplePerson{
		Name: "AAND",
		Age:  14,
	}
	multisortExamplePersons := []multiSortExamplePerson{p1, p2, p3, p4}

	return multisortExamplePersons
}

// ExampleMultiSorted_multiKey Function, sorts the input with two keys
// Sorts slice of multiSortExamplePerson with key: Name in Descending Order and then with Key: Age in Ascending
// Prints the output, after converting the []interface back to expected(required) type
func ExampleMultiSorted_multiKey() {
	multisortExamplePersons := makeInput()
	sortKeys := []string{"Name", "Age"}
	ascendingOrder := []bool{false, true}
	multiSortResponse, err := MultiSorted(multisortExamplePersons, sortKeys, ascendingOrder)
	if err != nil {
		return
	}
	for i := range multiSortResponse {
		multiSortResponse[i] = multiSortResponse[i].(multiSortExamplePerson)
	}
	fmt.Println(multiSortResponse)
	// Output: [{Joe 26} {Azin 14} {AAND 14} {Bold 11}]
}

// ExampleMultiSorted_singleKey1 Function, sorts the input with one key
// Sorts slice of multiSortExamplePerson with key: Name in Ascending Order
// Prints the output, after converting the []interface back to expected(required) type
func ExampleMultiSorted_singleKey1() {
	multisortExamplePersons := makeInput()
	sortKeys := []string{"Name"}
	ascendingOrder := []bool{true}
	multiSortResponse, err := MultiSorted(multisortExamplePersons, sortKeys, ascendingOrder)
	if err != nil {
		return
	}
	for i := range multiSortResponse {
		multiSortResponse[i] = multiSortResponse[i].(multiSortExamplePerson)
	}
	fmt.Println(multiSortResponse)
	// Output: [{Joe 26} {Bold 11} {Azin 14} {AAND 14}]
}

// ExampleMultiSorted_singleKey2 Function, sorts the input with one key
// Demonstrates that the default ascending order is applied, when nothing is provided with ascendingOrder slice input
// Sorts slice of multiSortExamplePerson with key: Age in Ascending Order (Taken as default)
// Prints the output, after converting the []interface back to expected(required) type
func ExampleMultiSorted_singleKey2() {
	multisortExamplePersons := makeInput()
	sortKeys := []string{"Age"}
	var ascendingOrder []bool
	multiSortResponse, err := MultiSorted(multisortExamplePersons, sortKeys, ascendingOrder)
	if err != nil {
		return
	}
	for i := range multiSortResponse {
		multiSortResponse[i] = multiSortResponse[i].(multiSortExamplePerson)
	}
	fmt.Println(multiSortResponse)
	// Output: [{Joe 26} {Azin 14} {AAND 14} {Bold 11}]
}

// ExampleHelp provides, a use case of how to convert the T.([]interface) into the required type
// desiredType is the type which the client sends in the data
func ExampleHelp() {
	fmt.Println(Help())
	// Output:
	/*
		outputSlice, err := MultiSorted(inputSlice, inputKeys, inputOrder)
	for i := range outputSlice {
		outputSlice[i] = outputSlice[i].(desiredType)
	}
	*/
}
// TestMultiSorted, performs table testing with All possible combination of 2 keys and 2 orders. Taking each key one at a time and in combination
func TestMultiSorted(t *testing.T) {
	// variable declaration for Type multiSortExamplePerson to be tested as a slice
	p1 := multiSortExamplePerson{
		Name: "Joe",
		Age:  26,
	}
	p2 := multiSortExamplePerson{
		Name: "Azin",
		Age:  14,
	}
	p3 := multiSortExamplePerson{
		Name: "Bold",
		Age:  11,
	}
	p4 := multiSortExamplePerson{
		Name: "AAND",
		Age:  14,
	}

	// Input matrix for keys to be sorted with
	sortKeysmatrix := [][]string{
		{"Name"},
		{"Name", "Age"},
		{"Age"},
		{"Name"},
		{"Name", "Age"},
		{"Age"},
		{"Name", "Age"},
		{"Name", "Age"},
	}

	// Input matrix for the order in which the slice must be sorted
	ascendingOrderSlice := [][]bool{
		{true},
		{true, true},
		{true},
		{false},
		{false, false},
		{false},
		{true, false},
		{false, true},
	}

	// Expected results based on above mentioned key and keyOrder combination
	expectedResults := [][]multiSortExamplePerson{
		{p4, p2, p3, p1},
		{p1, p4, p2, p3},
		{p1, p2, p4, p3},
		{p1, p3, p2, p4},
		{p1, p2, p4, p3},
		{p1, p2, p4, p3},
		{p1, p2, p4, p3},
		{p1, p2, p4, p3},
	}

	// For all the expectedResults
	for i, v := range expectedResults {
		// create the input
		multisortExamplePersons := makeInput()
		// send it to the Multi Sort, to sort
		multiSortResponse, err := MultiSorted(multisortExamplePersons, sortKeysmatrix[i], ascendingOrderSlice[i])
		if err != nil {
			return
		}
		// Convert the result and expectedResult to a string for easier comparison without the need of reflect package
		expectedResult := fmt.Sprint(v)
		multiSortResult := fmt.Sprint(multiSortResponse)
		// Check for the match, if the expectedResult does not match multiSortResult, fail the test
		if expectedResult != multiSortResult {
			t.Errorf("Expected: %v Got %v", expectedResult, multiSortResult)
		}
	}
}

// TestHelp to check, if Help function remains unchanged
func TestHelp(t *testing.T) {
	expectedResult := `outputSlice, err := MultiSorted(inputSlice, inputKeys, inputOrder)
	for i := range outputSlice {
		outputSlice[i] = outputSlice[i].(desiredType)
	}`
	actualResult := Help()
	if expectedResult != actualResult {
		t.Errorf("Expected: %v Got %v", expectedResult, actualResult)
	}
}
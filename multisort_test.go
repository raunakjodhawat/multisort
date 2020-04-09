package multisort

import (
	"fmt"
	"testing"
)

type multiSortExamlePerson struct {
	Name string
	Age  int
}

func makeInput() []multiSortExamlePerson {
	p1 := multiSortExamlePerson{
		Name: "Joe",
		Age:  26,
	}
	p2 := multiSortExamlePerson{
		Name: "Azin",
		Age:  14,
	}
	p3 := multiSortExamlePerson{
		Name: "Bold",
		Age:  11,
	}
	p4 := multiSortExamlePerson{
		Name: "AAND",
		Age:  14,
	}
	multisortExamplePersons := []multiSortExamlePerson{p1, p2, p3, p4}

	return multisortExamplePersons
}

// Examples
func ExampleMultiSort_MultiKey() {
	multisortExamplePersons := makeInput()
	sortKeys := []string{"Name", "Age"}
	ascendingOrder := []bool{false, true}
	multiSortResponse, err := MultiSort(multisortExamplePersons, sortKeys, ascendingOrder)
	if err != nil {
		return
	}
	for i := range multiSortResponse {
		multiSortResponse[i] = multiSortResponse[i].(multiSortExamlePerson)
	}
	fmt.Println(multiSortResponse)
	// Output: [{Joe 26} {Azin 14} {AAND 14} {Bold 11}]
}
func ExampleMultiSort_SingleKeyType_1() {
	multisortExamplePersons := makeInput()
	sortKeys := []string{"Name"}
	ascendingOrder := []bool{true}
	multiSortResponse, err := MultiSort(multisortExamplePersons, sortKeys, ascendingOrder)
	if err != nil {
		return
	}
	for i := range multiSortResponse {
		multiSortResponse[i] = multiSortResponse[i].(multiSortExamlePerson)
	}
	fmt.Println(multiSortResponse)
	// Output: [{Joe 26} {Bold 11} {Azin 14} {AAND 14}]
}

func ExampleMultiSort_SingleKeyType_2() {
	multisortExamplePersons := makeInput()
	sortKeys := []string{"Age"}
	var ascendingOrder []bool
	multiSortResponse, err := MultiSort(multisortExamplePersons, sortKeys, ascendingOrder)
	if err != nil {
		return
	}
	for i := range multiSortResponse {
		multiSortResponse[i] = multiSortResponse[i].(multiSortExamlePerson)
	}
	fmt.Println(multiSortResponse)
	// Output: [{Joe 26} {Azin 14} {AAND 14} {Bold 11}]
}

// tests
func TestMultiSort(t *testing.T) {
	p1 := multiSortExamlePerson{
		Name: "Joe",
		Age:  26,
	}
	p2 := multiSortExamlePerson{
		Name: "Azin",
		Age:  14,
	}
	p3 := multiSortExamlePerson{
		Name: "Bold",
		Age:  11,
	}
	p4 := multiSortExamlePerson{
		Name: "AAND",
		Age:  14,
	}

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
	expectedResults := [][]multiSortExamlePerson{
		{p4, p2, p3, p1},
		{p1, p4, p2, p3},
		{p1, p2, p4, p3},
		{p1, p3, p2, p4},
		{p1, p2, p4, p3},
		{p1, p2, p4, p3},
		{p1, p2, p4, p3},
		{p1, p2, p4, p3},
	}
	for i, v := range expectedResults {
		multisortExamplePersons := makeInput()
		multiSortResponse, err := MultiSort(multisortExamplePersons, sortKeysmatrix[i], ascendingOrderSlice[i])
		if err != nil {
			return
		}
		a := fmt.Sprint(v)
		b := fmt.Sprint(multiSortResponse)
		if a != b {
			t.Errorf("Expected: %v Got %v", a, b)
		}
		fmt.Println(v, multiSortResponse, a == b)

	}
}

## Getting Started
### Purpose
This repo aims to increase the functionality of default Sort package provided by Golang.
It's a extremely light weight program, that exports just two functions:
1. multisort(inputSlice, inputKeys, inputOrders) []interface{}
2. Help()

### Installation
```
go get -d github.com/raunakjodhawat/multisort
```
### Usage
```
    sortedSlice, err := multisort.MultiSort(inputSlice, inputKeys, SortOrders)

    fmt.Println(Help())
```

#### Details
##### Function - MultiSort
    - Three Input Parameter
        -- inputSlice (The Slice you want to Sort, can be of Struct of any form)
            --- Currently, the program can not sort Slices with the underlying struct's containing non-primitive type
        -- inputKeys (Slice of String based keys, by which you want to sort)
            --- Make sure that Keys match, Capitilization is important with what's defined in the struct
        -- SortOrders (Slice of Booleans, indicating the sorting direction)
            --- default sort orders is Ascending
    - Two Output
        -- A slice of sorted Array
            --- Sorted array if required can be converted to user defined Struct, using the below function
                outputSlice, err := MultiSort(inputSlice, inputKeys, inputOrder)
                	for i := range outputSlice {
                		outputSlice[i] = outputSlice[i].(desiredType)
                	}
        -- Error, if present
        
```
Usage:
// Define input parameters  
    sortKeys := []string{"Name", "Age"}
    ascendingOrder := []bool{false, true}

// Send it to multisort function
    multiSortResponse, err := multisort.MultiSort(persons, sortKeys, ascendingOrder)
    
    if err != nil {
        // return/ Print /Panic
    }
// Traverse the function, and convert each element back to the desired type    
    for i := range multiSortResponse {
        persons[i] = multiSortResponse[i].(Person)
    }
// Use the sorted Slice
    fmt.Println(persons)
```

##### Function - Help
    - No Input, More of a utility function which returns a String helping the
        developer figure out how to handle the interface object and convert it
        to the desired Type T
```
Usage:
    multisort.Help()
```

[![GoDoc](https://godoc.org/github.com/raunakjodhawat/multisort?status.svg)](https://godoc.org/github.com/raunakjodhawat/multisort)  

[Website](https://raunakjodhawat.github.io/multisort/)

[LinkedIn](https://www.linkedin.com/in/jodhawat/)
##### External Dependencies
1. fmt
2. reflect
3. testing
4. Sort
5. Sync

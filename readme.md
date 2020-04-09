# Purpose
This repo aims to increase the functionality of default Sort package provided by Golang

# Installation
```
go get -d github.com/raunakjodhawat/multisort
```
# Usage
```
// Define input parameters  
    sortKeys := []string{"Name", "Age"}
    ascendingOrder := []bool{false, true}

// Send it to multisort function
    multiSortResponse, _ := multisort.MultiSort(persons, sortKeys, ascendingOrder)
// Traverse the function, and convert each element back to the desired type    
    for i := range multiSortResponse {
        persons[i] = multiSortResponse[i].(Person)
    }
// Use the sorted Slice
    fmt.Println(persons)

    help()
```
## External Dependency
1. fmt
2. time
3. reflect
4. testing


### Limitation
1. us
### Website
```
    https://raunakjodhawat.github.io/multisort/
```

[![GoDoc](https://godoc.org/github.com/raunakjodhawat/multisort?status.svg)](https://godoc.org/github.com/raunakjodhawat/multisort)  

[Website](https://raunakjodhawat.github.io/multisort/)

inner type is always defined by go. I am using that

no fatal is send

contributuions
"%#V" use
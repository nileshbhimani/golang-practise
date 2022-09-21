
package main

import (
	"fmt"
    // "reflect"
    "sort"
)


func main() {
    var MapEX = map[int]string {1:"no",2:"nil"}

    fmt.Println(MapEX)


    var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)
 
	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)


    newList:= make(map[string]int)
    newList["vishu"] = 1
    newList["nil"] = 2
    fmt.Println(newList)
    fmt.Println(len(newList))

    fmt.Println(newList["vishu"],newList["nil"])

    delete(employee, "Mark")
    fmt.Println(employee)

     for key, element := range employee {
        fmt.Println("Key:", key, "=>", "Element:", element)
    }



    unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
 
	keys := make([]string, 0, len(unSortedMap))
 
	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

    fmt.Println("Key values",keys)
 
	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}


    values := make([]int, 0, len(unSortedMap))
 
	for _, v := range unSortedMap {
		values = append(values, v)
	}
 
 // Sort slice values.
	sort.Ints(values)
 
 // Print values of sorted Slice.
	for _, v := range values {
		fmt.Println(v)
	}

}

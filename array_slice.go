
package main

import (
	"fmt"
    "reflect"
)

func main() {
        var array []int
        var stringarray []string

        nslice := make([]int,0,2)
            for i:=0;i<=10;i++{
                 nslice = append(nslice,i)   
                fmt.Println("length",len(nslice))
                fmt.Println("length",cap(nslice))    
            }

            fmt.Println("Array slice",nslice)
        

        var NoMake = []int{1,2,3,4}
        // var NoMakeString = []int{"go","no"}

        for index,value:= range NoMake{
            fmt.Println(index,"====>",value)
        }


        var newKeyWord = new([50]int)[0:10]

            newKeyWord = append(newKeyWord,45)    

           fmt.Println("new key word append",newKeyWord)

        fmt.Println(reflect.ValueOf(array).Kind())
        fmt.Println(reflect.ValueOf(stringarray).Kind())

}


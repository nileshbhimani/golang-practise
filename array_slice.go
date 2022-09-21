
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


        a := make([]int, 2, 5)
        a[0] = 10
        a[1] = 20

        a = append(a,30,40,50)
        fmt.Println("value of a",a)

        var b = make([]int,8,16)

        copy(b,a)

        fmt.Println("Copy of a to b == b",b)


        var strArr = []string {"ni","nil"}

        strArr = RemoveIndex(strArr,0)
         fmt.Println("Removed value of a",strArr)

        fmt.Println(reflect.ValueOf(array).Kind())
        fmt.Println(reflect.ValueOf(stringarray).Kind())


       var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
        fmt.Println(itemExists(strSlice, "Canada"))
        fmt.Println(itemExists(strSlice, "Africa")) 

}

func itemExists( Arr interface{}, t interface{}) bool{
    s := reflect.ValueOf(Arr)

    fmt.Println("Value of if Exist check",s)

    // for index,value := range Arr{
    //     if (value == t){
    //         return true    
    //     }
    // }    

    for i:=0;i<s.Len();i++{
        if ( s.Index(i).Interface()==t){
            return true
        } 
    }

    return false
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

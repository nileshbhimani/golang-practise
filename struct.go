
package main

import (
	"fmt"
    // "reflect"
    // "sort"
    "encoding/json"    
)

/// https://www.golangprograms.com/go-language/struct.html

type rectangle struct{
    length  int
	breadth int
	color   string

    geometry struct {
		area      int
		perimeter int
	}
}


type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"
}

// nested usage of Salary into Employee
// type Salary struct {
// 	Basic, HRA, TA float64
// }

// type Employee struct {
// 	FirstName, LastName, Email string
// 	Age                        int
// 	MonthlySalary              []Salary
// }

type Emp struct{
    Fname string `json:"firstname"`
    Lname string `json:"lastname"`
}

func main() {
    // s := rectangle{10.5,6.5,"RED"}

    var rect rectangle // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

    rect.geometry.area = rect.length * rect.breadth
    rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

    fmt.Println(rect)
    fmt.Println(rect.geometry)

    rect4 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(rect4)

    rect1 := new(rectangle) // rect1 is a pointer to an instance of rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)

    // var rect2 = &rectangle{10, 20, "Green"} // Can't skip any value
	// fmt.Println(rect2)

     jStr:=  `{
                    "firstname": "Rocky",
                    "lastname": "Sting"
                }`

    e1:= new(Emp)

    json.Unmarshal([]byte(jStr), e1)
    fmt.Println(e1)

    e2:=new(Emp)
    e2.Fname = "Ramesh"
    e2.Lname = "Soni"
  
    jsonStr, _ := json.Marshal(e2)
    fmt.Printf("%s\n", jsonStr)
    
    e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	fmt.Println(e.EmpInfo())

}

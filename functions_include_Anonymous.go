package main

import "fmt"

func Adding(a int, b int) (sum int, sub int){
     sum = a + b
     sub = a - b   

    return
}

func Update(t *int,y *string){
    *t = *t + 5      // defrencing pointer address
	*y = *y + " Doe" // defrencing pointer address
	return
}

func main(){
    var a, p int
	a, p = Adding(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)  

    var x,z = 50,"nil"
    fmt.Println("before",x,z)
    Update(&x,&z)
    fmt.Println("After",x,z)


}


================================= Anonymous ====================================

package main

import "fmt"

var (
    area = func(len int,height int) int{
        return len * height
    }
)

func main(){
    fmt.Println(area(20,30))        
}

========================== User Defined type called with nested structure =======================
package main

import "fmt"

type First func(int) int
type Second func(int) First

func squareSum(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))
}

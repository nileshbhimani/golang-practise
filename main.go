package main


import (
	"fmt"
)

type info struct{
    UserId int
    Password string
}

//    var UserId int
//    var Password string

func login() int{
    s:=info{}

    if ( s.UserId != 12 && s.Password != "2345" ) {
        fmt.Println("please enter username \n")	
          fmt.Scan(&s.UserId)
        fmt.Println("please enter password \n")
          fmt.Scan(&s.Password)  
    }	

    if( s.UserId == 12 && s.Password == "2345" ){
        return 1
    }
    
    return 0
}

func CreateAccount(){
    c:=info{}
    fmt.Println("please enter username \n")	
        fmt.Scan(&c.UserId)
    fmt.Println("please enter password \n")
        fmt.Scan(&c.Password)

    fmt.Println("Account is created succesffuly with","user id ",c.UserId,"Password",c.Password)
}

func main(){
    
    i:= login()
    

    if( i== 0 ){
        ll:= 0
        fmt.Println("Invalid credential, Please select option to continue") 
        fmt.Println("1.Login")
        fmt.Println("2.Create account")
        fmt.Println("3.Quite")    
        fmt.Scan(&ll)
        switch ll {
            case 1:
                fmt.Println("Login")
                i=login()
            case 2:
                fmt.Println("Create account")
                CreateAccount()
            case 3:
                break
            default:
                fmt.Println("Please select option to continue")
	    }
        
        // login()
    }else{
        fmt.Println("Hi! Welcome to Mr. Rohit ATM Machine!")

        ll:= 0
        fmt.Println("Please select below option to continue") 
        fmt.Println("1.Withdraw money")
        fmt.Println("2.Deposit money")
        fmt.Println("3.Request balance")   
        fmt.Println("Quite")    
        fmt.Scan(&ll)
        var wd int
        balance := 10000
        switch ll {
            case 1:
                fmt.Println("Please enter Withdraw money from balance",balance)
                fmt.Scan(&wd)
                if(wd!=0){
                     fmt.Println("Money is successfully withdraw",wd)
                     fmt.Println("remaining balance is",balance - wd)
                }
               
            case 2:
                fmt.Println("Please enter Deposit money")
                fmt.Scan(&wd)
                if(wd!=0){
                     fmt.Println("Money is deposited succesfully",wd)
                     fmt.Println("remaining balance is",balance + wd)
                }
            case 3:
                 fmt.Println("Your balance is",balance)   
            case 4:
                break
	    }
    
    } 

 }	

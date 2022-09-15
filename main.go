package main


import (
	"fmt"
)

   var UserId int
   var Password string
	
	   fmt.Println("please enter username");	
	   UserId = fmt.Scan(&UserId)
	   fmt.Println("please enter password");
	   Password = fmt.Scan(&UserId)	

	   fmt.Println("your userid and password is", UserId, Password);		

//    if ( userId == 12 &&  password == 2345 ) {
// 		fmt.Println("plese enter c to stop");
// 		fmt.Scan(&com)
//    }else{
// 		fmt.Println("plese enter c to stop");
// 		fmt.Scan(&com)
//    } 

//     for com != "c"{
//         fmt.Println("plese enter c to stop");
//         fmt.Scan(&com)
//     } 
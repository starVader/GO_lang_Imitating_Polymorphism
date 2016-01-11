package main 

import "fmt"

func printFirstname(firstname interface{}, surname  ...interface{}) (interface{}, interface{}){
	unmaskedName := surname									// we cannot have variadic returns
	if unmaskedName != nil{
		return firstname, surname
	
	}else {
		return firstname,""
	}
}

func main(){
	fmt.Println(printFirstname("joe","Blow","clown"))
	fmt.Println(printFirstname("joe"))
}

/* swift equivalent
func printFirstname(firstname:string, surname:string?=nil) {
	if let (unmaskedName = surname){
		Print(firstname, surname[0])
	
	}else {
		Print(firstname)
	}
}
printFirstname("joe","Blow"
)printFirstname("joe")
*/

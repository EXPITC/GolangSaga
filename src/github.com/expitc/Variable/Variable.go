package main

import (
	"fmt"
)

// you can declare variable at package level
// you can't use :=  (semicolon equal syntax) you must have full declaration syntax
var num int = 32

// we can do like this
var (
	totalpatient int = 12
	doctorNumber int = 3
	hospitalName string =  "IJY"
)
/* rather than doing this :
	var totalpatient int = 12
	var doctorNumber int = 3
	var hospitalName string = "IJY"
 we can wrap it to make it more clean
*/

// name convention on Go and 3 visibility
var i int = 1 //this name with lowercase will just only available on "this" package lv which mean anything that consume "this" package can't consume or can't see it
var I int = 2 //this name with uppercase will expose to outside world which mean the opposite of lowercase , this uppercase variable will export at the front of package and it's globally visible and can be consume by other packages


var shadowing string = "you can declare on this scope"
func main() {
	var block string = "block lv(func main(){} ) visibility never be expose to outside"
	fmt.Println(block)
	var shadowing string = "but it still works fine its okay to declare on this scope or this way its call shadowing"
	fmt.Println(shadowing);
	// go must be declare type just like typescript
	// var i int = 42;
	// but go can figure it out what type to assign with this way :=
	// you can say its shorthand but it's depends on condition or context
	i := 42
	whenGoDontHaveInformationToAssignUse := 99.
	var whenYouHaveInformationToAssignTypeUse float32 = 00.1
	fmt.Println(i)
	fmt.Printf(
		"%v, %T",
		whenGoDontHaveInformationToAssignUse,
		whenGoDontHaveInformationToAssignUse,
	)
	fmt.Printf("| when you have information to assign the type use this := %T", whenYouHaveInformationToAssignTypeUse)
}

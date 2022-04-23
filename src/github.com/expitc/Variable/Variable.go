package main

import (
	"fmt"
)

func main() {
	// go must be declare type just like typescript
	// var i int = 42;
	// but go can figure it out with this way :=
	// you can say its shorthand but it's depends on condition or context
	i := 42
	whenGoDontHaveInformationToAssignUse := 99
	fmt.Println(i)
	fmt.Printf(
		"%v, %T",
		whenGoDontHaveInformationToAssignUse,
		whenGoDontHaveInformationToAssignUse,
	)
}

package main

import (
	"fmt"
	"strconv" // for conversion to string format
)
func main() {
	var i int = 42
	fmt.Printf("%v, %T", i ,i)
	
	var j float32
	j = float32(i)
	// this how go conversion the type but be carefull
	fmt.Printf("%v, %T", j ,j)
	
		// Carefull
	var o float32 = 42.5
	fmt.Printf("%v, %T", o, o)

	var p int
	p = int(o) // be carefull of this the first e.g works well because int to float
	/* but when you change float to int float loss his ability on int and cause
	loss information the o value should be 42.5 and when you convert that float to int then 
	float will lose his ability and p became 42.
	*/
	fmt.Printf("%v, %T", p ,p)

	/* if you want to change int to string or you want 42 to string you can't just do it like this:
	 var a string
	a = string(i) : output = *
	that because in golang they recognize string as alias of byte
	so you need o import "strconv"
	*/
	var a string
	// Itoa "I" stand for integer "to" mean to and "a" mean ASCII 
	// Itoa same with strconv.FormatInt(int64(x), 10)
	a = strconv.Itoa(i)
	a = strconv.FormatFloat(float64(o), 'f', 6, 64) // you can do it conversion to float64(o) since FormatFloat only accept 64 bit or just change o to float64
	fmt.Printf("%v, %T", a ,a)
}

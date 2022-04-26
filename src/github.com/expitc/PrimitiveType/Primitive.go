package main

import (
	"fmt"
)

func main() {
	//[Boolean]
	var n bool = false
	fmt.Printf("%v,%T\n", n, n)
	// this is boolean too
	n1 := 1 == 1
	n2 := 1 == 2
	fmt.Printf("%v,%T\n", n1, n1)
	fmt.Printf("%v,%T", n2, n2)

	// in go all variable have default values and the default values was 0 , 0 mean false on boolean
	var defaultVal bool
	// it will print false
	fmt.Printf("%v,%T\n", defaultVal, defaultVal)

	//[integer]
	in := 42 //this type int
	/* but in go we have another type
	int8  -128 - 127
	int16 -32 768 - 32 767
	int32  and so on
	int64 and so on
	*/
	var uin uint = 12
	/*
		but have unsigned integer or uint
		uint8 0 - 255
		uint16 0 - 65 535
		uint32 and so on
		but we don't have 64 of unsigned integer
	*/
	fmt.Printf("%v,%T\n", in, in)
	fmt.Printf("%v,%T\n", uin, uin)

	a := 10 //1010
	b := 3  //0011
	// just like we doing division we can't change the type
	// so dividing integer and integer can't give us a float point number, for example
	// we're also not allowed to add integer with diffrent type
	// var a int = 32 , var b int8 = 1 , we can't (a + b) , you need same type so do like (a + int(b))
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	// bitwise
	fmt.Println(a & b)  // and operator so bit will look like 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100
	// bitshifting
	a = 8               //2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1

	//[floating point number]
	// in go float point number follow III E 754 and in that standard we're going to pull out 32 bit float number and 64 bit float number
	// float32 ±1.18E-38 - ±3.4E38
	// float64 ±2.23E-308 - ±1.8E308
	// declaration float
	f := 3.14 //when you declare with initializer syntax like this it's always float64
	// var n float32 = 3.14 // then 13.7e72 is to big but the point is you can write float like
	f = 13.7e72 //this
	f = 2.1e14  //on go
	fmt.Printf("%v,%T", f, f)

	fl := 10.2
	fll := 3.7
	// return or result always be floating
	// reminder(%) and bitwise or bitshifting operation only available for int
	fmt.Println(fl + fll)
	fmt.Println(fl - fll)
	fmt.Println(fl * fll)
	fmt.Println(fl / fll)

	// complex number
	// 2 type of complex number , complex64 and complex128
	// basically reason behind that is we taking float32 plus a float32 or a float64 and float64
	// var compA complex64 = 1 + 2i
	compA := 1 + 2i
	compB := 2 + 5.2i
	fmt.Println(compA + compB)
	fmt.Println(compA - compB)
	fmt.Println(compA * compB)
	fmt.Println(compA / compB)
	// convert to real and imaginary number
	// you do real or imaginary with complex64 then it will give you float32 and if you do complex128 then float64
	fmt.Printf("%v,%T", real(compA), real(compA))
	fmt.Printf("%v,%T", imag(compA), imag(compA))
	// convert to complex
	compA = complex(5,12) //the first number is real part and the second is imaginary part

	//[texting type]
	// string
	// string stand for utf-8 it's powerfull but can't encode every type of character that's available
	s := "this is a string"	// we can actually treat this string like short collection on array similar with python. its collection of string
	fmt.Printf("%v,%T\n", s[2], s[2])
	// the result will be 105, uint8 why? because string is an alias for bytes on go
	// so if you want to get the string back then convert with string(s[2])
	fmt.Printf("%v,%T\n", string(s[2]), s[2])
	s2 := "and this a string too"
	fmt.Println(s + s2) //of course concat works fine
	// we can do convertion string to collection of byte
	bs:= []byte(s)
	fmt.Printf("%v,%T\n", bs, bs)
	
	// rune
	// rune stand for UTF8 an a alias of true type alias for int32
	// rune declare with '' and string with double quotes ""
	var r  rune = 'a'
	fmt.Printf("%v,%T\n", r, r) // result 97, int32
	// if you need to work with UTF8 you can look at doc at reader.ReadRune
	// they have ReadByte , ReadRune etc
}

package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

type ExtendedFloat float64

func (e ExtendedFloat) toString() string {
	return strconv.FormatFloat(float64(e), 'f', -1, 64)
}

func IntegerDT() {
	// available int types
	var a int8 = 127                  // Range for int8: -128 to 127, 1 bit for sign and 7 bits for magnitude of number, 2^7 = 128
	var _ int16 = 32767               // Range for int16: -32,768 to 32,767
	var _ int32 = 2147483647          // Range for int32: -2,147,483,648 to 2,147,483,647
	var _ int64 = 9223372036854775807 // Range for int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

	// max values using math functions
	var _ int32 = math.MaxInt32

	var _ uint8 = 255                  // Range for uint8: 0 to 255
	var _ uint16 = 32767               // Range for int16: -32,768 to 32,767
	var _ uint32 = 2147483647          // Range for int32: -2,147,483,648 to 2,147,483,647
	var _ uint64 = 9223372036854775807 // Range for int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

	// natural and most efficient, size depends on system, 32bit/64 bit
	var _ int = 9223372036854775807  // -2^31 to 2^31-1 on 32-bit systems and from -2^63 to 2^63-1 on 64-bit systems
	var _ uint = 9223372036854775808 // 0 to 2^32-1 on 32-bit systems and from 0 to 2^64-1 on 64-bit systems

	var _ rune = 2147483647 // rune is synonym to int32, holds unicode values

	// default/implicit value
	//f := 273456 // this will have an implicit type of int(natural int)

	var _ uintptr = 78689689 // holds bits of a pointer value

	// Explicit type conversion

	// int, uint, int8, int16 etc. are all DIFFERENT types and will need explicit type conversions
	// var o int16 = a  //this is invalid
	var _ int16 = int16(a) //this is valid
}

func FloatDT() {
	// available float type
	var _ float32 = math.MaxFloat32 // Range: Approximately -3.4 x 10^38 to 3.4 x 10^38, Precision: About 7 decimal digits
	var _ float64 = math.MaxFloat64 // Range: Approximately -1.8 x 10^308 to 1.8 x 10^308, Precision: About 15 decimal digits

	// default/implicit value
	//f := 273456.2389745823 // this will have an implicit type of float64
}

func ComplexDT() {
	// available complex type
	var _ complex64 = complex(1, 2)   // 1+2i, Real part and imaginary part are both float32
	var _ complex128 = complex(4, 78) // 4+78i, Real part and imaginary part are both float64
}

func BooleanDT() {
	var _ bool = true
	var _ bool = false
}

func StringDT() {
	// In go a string is an immutable sequence of bytes
	// Go uses UTF-8 encoding for stings
	// rune is used to work with Unicode characters and text, whereas byte is used for working with binary data or ASCII characters.

	// USE RUNE when working with string-characters.

	// 1. String to char array([]rune)
	str := "Hello, 世界" // string literal
	runeArray := []rune(str)

	//2. Iterate over char array([]rune)

	// Iterate over the array of runes and print each character
	for _, char := range runeArray {
		fmt.Println(char) // prints the int value of the character
	}

	// 3. Pick char of a string
	fmt.Println(int(str[0])) // 72, str[0] is of type uint8

	// 4. Find length function (ASCII) | returns number of bytes in a string
	fmt.Println(len(str))

	// Conversion

	// 5. Int to Char
	fmt.Printf("Integer %d corresponds to character: %c\n", str[0], rune(str[0]))

	// 6. Int to String
	fmt.Printf("Integer %d corresponds to string: %s\n", str[0], string(rune(str[0])))

	// 7. Find the length of the string Unicode+ASCII (number of runes) - ALWAYS USE THIS
	length := len([]rune(str))
	fmt.Printf("Length of the string: %d\n", length)
	// or
	length = utf8.RuneCountInString(str)
	fmt.Println(length)

	// []byte and []rune to string
	s := string(runeArray)
	fmt.Printf("byte to string conversion: %s", s)

	// STRING MANIPULATION

	/* packages:
	1. bytes: Because strings are immutable, building up strings incrementally can involve a lot of allocation and copying.
	in such cases, it's more efficient to use the bytes.Buffer type.
	2. strconv: provides functions for converting boolean, integer, and float values to and from string.
	3. unicode: provides functions like IsDigit, IsLetter, IsUpper, IsLower for classifying runes.*/

	// String as an array of Bytes: Immutable
	str = "Hello, World!"
	_ = str[0] // Access the byte at index 0 ('H'), Attempting to modify a string by assigning a new value to a specific
	// byte index will result in an error.

	// String as a Slice of Bytes: Mutable
	byteSlice := []byte(str)
	byteSlice[0] = 'C' // Modifying the byte at index 0

	_, err := strconv.Atoi("123")            // string to int
	_, err = strconv.ParseInt("123", 10, 64) // string to int
	_, err = strconv.ParseFloat(str, 64)     // string to float
	_, err = strconv.ParseBool("true")       // string to bool

	_ = strconv.FormatFloat(56.78, 'f', -1, 64) // float to string
	_ = strconv.Itoa(56)                        // int to string
	_ = strconv.FormatBool(false)               // bool to string

	// Constants
	// evaluation at compile time
	const pi = 3.14159

	fmt.Println(err)
}

## Program Structure -  
   * ***[Names](#names--)*** 
   * ***[Declerations](#declerations--)***
   * ***[Variables](#variables--)***


## Basic Data Types -  
   * ***[Integers](#integers--)*** 
   * ***[Floating Point Numbers](#floating-point-numbers--)*** 
   * ***[Complex Number](#complex-number--)*** 
   * ***[String](#string--)*** 


### Names -
----------

- names are case sensitive.
- if entity declared inside a function, its local to that function.
- ***if declared outside of a func then it is visible in all files of the package it belongs*** (kind of package private).
- ***the case of the first letter of the name*** determines its visibility accross package boundaries. i.e. if ***name begins with uppercase then it is **exported** which means visible and accessible outside its package and maybe referred to by other parts of the program*** , e.g. Printf in fmt package , Stdin in os package.



### Declerations -
----------

- decleration names a program entity and also specifies some or all of its properties.
- a func decleration can has a ***optional list of results***, result list is omitted if the function does not have a return type.

- sample function -
```go
    package main
    import (
        "fmt"
    )

    func main() {
        const freezingF , boilingF = 32.0 , 212.0
        fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
        fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
    }

    func fToC(temp float64) float64 {
        return (temp - 32) * 5/9
    }
```


### Variables -
----------




 









### Integers -
----------


- numeric datatype has serveral variation of int, float and complex number which determines their ***sizes and signedness*** i.e GO has both signed and unsigned arithmetic.
- for ints there are 4 variation of size i.e. int8, int16, int32 and int64. corresponding unsigned are uint8, uint16 etc.. ***but there are also just int and uint that natural and most efficient size for signed and unsigned in a particular platform***(size of these can vary based on compiler 32 or 64 bit even on identical hardware) 
- **type rune** is a synonym of int32 and by convention it indicates value of a unicode codepoint. here 2 names can be used interchangably like byte is uint8 but the former emphasizes that it is a raw data rather than just a small number.
- **uintptr** a specific unsigned integer type that's size is not specified but good enough to hold all bits of a pointer value. (normally used in low level programming i.e. boundary of a GO program with C)
- explicit conversion is required for assigned int32 or int64 value to int or any other type.
- signed number are represented in 2's complement, with range from −2^n−1 to (2^n−1)−1 , for a n bit number
- GO's binary operator for arithmetic, logic and comparision precendence order below - 
![alt](/resources/images/Go_operator_precedence.PNG)

- ***remainder operator % only applies to integer*** and sign of remainder depends on sign of divident.
- GO also provides bitwise operator, the first 4 below treat their operands as bit patterns with no concept of carry or sign - XOR behaves xor in binary operation but negation or complement operator in unary operation. 
![alt](/resources/images/Go_bitwise_operator.PNG)


- arithmetically left shift x<< n is equivalent to multiplication by 2^n and right shift x >> n equivalent to floor of division by 2^n
- unsigned numbers tend to be used only when their bitwise operators or peculiar arithmetic operators are required,as when implementing bit sets, parsing binary file formats,or for hashing and cryptography.They are typically not used for merely non-negative quantities.

- for every type T, conversion operator T(x) converts x to type T if conversion is allowed. ***e.g -  x int16 = 3 ; y int = int(x)***
- integer number can be either written as normal number or octal starting with 0 or hexadecimal starting with 0x or 0X(small or capital case does not matter)
- octal numbers nowadays are only used for ***file permission on POSIX system*** , but hexadecimal are used widely to emphasize the bit pattern of a number over its numerical value
- when printing number with fmt package we can control radix and format using %d, %o and %x package.




### Floating point numbers -
----------

- they are of 2 sizes float32 and float64
- their range can be found in 'math' package, math.MaxFloat32 and math.MaxFloat64
- math package also provides utility like math.NaN (for things like infinity) and math.IsNaN() to check if a value is NaN (note NaN == NaN is false i.e NaN comparisons are not as expected)
- **do the float func to implement geometry function ??**


### Complex numbers -
----------

- there are 2 sizes of complex numbers complex64 and complex128 whose complements are float32 and float64
- complex(1,2) create a complex type with 1 as real and 2 as imaginary part (1+2i)
- build in real() and imag() are used to extract those real and imaginary parts of a complex number.
- math/cmplx pakcage provides library function for complex numbers like complex square root and exponential function
- **generate a mandelbrot fractel image**???


### String -
----------

- it is immutable sequence of bytes.
   - immutability means it is safe for 2 copies of string to share same memory, also a substring s[7:] shares the same memory so substring operation is also cheap.
- text strings are conventionally interpreted as utf8 encoded sequence of unicode code points(runes).
- len(string) gives number of bytes in string and not runes and s[i] retrieves ith byte i.e. s[1] will return int byte rather than char.
- attempting to access string outside its range say s[len(s)] results in a ***panic***.
- ith byte of a string is not necessarily the ith character of a string, because UTF-8 encoding on non-ASCII requires 2 or more bytes.
- substring can be achieved using s[i:j]
- 2 strings can be compared using == or > .the comparision happens byte by byte so result is natural lexiographical order.
- there are 2 types of string literal representation.

   1. ***Double quoted*** -   here escape sequences(like with \ ) can be used to insert arbitrary byte values into string. because GO's source file is interpreted as UTF-8 we can insert unicode code points in string literals.

   2. ***Raw string literal*** - `raw string literals` using backquotes instead of double quotes.
   here no escape sequence are processed (contents are taken literally including backslashes and new lines). the only processing is carriage returns are deleted so the strings are same on all platforms.
   
      - used for writing regular expression, which normally have lot of backslashes
      - useful for HTML templates, json literals, command usage messages(which extends over lines)

- under unicode all characters in world's writing system, accents etc. are assigned one standard number called ***Unicode code point*** , in GO terminology its ***rune***.
- utf-8 defines code point for over 120,000 chars for over 100 langs and scripts.
- natural datatype to hold single rune is int32.
- in utf-32 each code point has a same size of 32 bits and hence it is a wastage at times since mostly readable english text like ASCII requires 8 bits.
- ***UTF-8 is a variable length***  encoding of Unicode code points in bytes(created by Ken Thompson and Rob Pike). it uses between 1 and 4 bytes to represent each rune. (high order bits of the first byte of encoding for rune indicate how many bytes to follow like high-order 0 indicates 7-bit ASCII here each rune only takes 1 byte, high-order 110 indicate rune takes 2 bytes.)
![alt](/resources/images/utf8encoding_table.PNG)

- ***unicode*** package provide fn for working with indivisual runes (like distinguishing letters from numbers or lower case from upper case etc..)
- ***unicode/utf8*** package provide fn for encoding(writing runes as byte to some byte[] etc.)/decoding(getting runes from a string and its size) runes as bytes using utf-8.
- unicode escape in GO string allows us to specify them by their numeric code point value e.g. \uhhhh (16 bit value) and \Uhhhhhhhh(32 bit value)
- ideally if we had to loop over string and get runes using "utf8.DecodeRuneAsString()" could be used but then we would have a clumsy for loop to skip index based on rune size, ***this can be handled better using GO's range loop***.
```go
    // normal but clumsy way
    s := "Hello, 世界 , हर कोई!"
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("unicode char is : %[1]c, code point value is: %[1]v , size is %[2]d \n", r, size)
		s = s[size:]
    }
    
    // better way using range , when applied over string, performs UTF-8 decoding implicitly
	for index, r := range s {
		fmt.Printf("unicode char is : %[1]q, code point value is: %[1]v , currnet index is : %[3]d \n", r, index)
	}

```

- []rune conversion when applied to string returns sequence of unicode code points that string encodes.
- ***coverting a interger value to a string interprets the integer as a rune value and yields utf8 representation e.g. string(65) = "A" and not "65"*** , if rune is invalid e.g. string(1234567) a replacement character '\uFFFD' �(white questionmark inside black hexagon) is replaced.

- useful packages when working and manipulating string are - 

   - ****strings**** - utilities for comparing, joining , splitting, searching , trimming, replacing string etc..
   - ****bytes**** - has similar funcs for manipulating slices of bytes , []byte. can be useful when building up string incrementally as normal process involves too much copying and allocation of memory due to string immutability in such cases bytes.Buffer could be used.
   - ****strconv**** - provides func for converting boolean, int, float values to and from their string representations, and also func for quoting and unquoting strings.
   - ****unicode**** - provides func like IsDigit, IsLetter, IsUpper etc for classifying runes.

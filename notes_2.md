## Program Structure -  
   * ***[Names](#names--)*** 
   * ***[Declerations](#declerations--)***
   * ***[Variables](#variables--)***


## Basic Data Types -  
   * ***[Integers](#integers--)*** 
   * ***[Floating Point Numbers](#floating-point-numbers--)*** 
   * ***[Complex Number](#complex-number--)*** 


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
- when printing number with fmt package we can control radix and format using %d, %o and %x package
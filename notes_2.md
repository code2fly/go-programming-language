## Program Structure -  
   * ***[Names](#names--)*** 
   * ***[Declerations](#declerations--)***

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
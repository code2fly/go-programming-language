# The GO programming language
=============

## Introduction -  
--------------

### Origins of GO -
Go influences from lot of languages like(C,pascal,CSP). image shows GO influence from its ancestors.

![alt text](go_influence)
[go_influence]: https://github.com/johnsnowtarg92/go-programming-language/resources/images/Go_origin.PNG "Go Influence tree"


#### Features inherited from different languages - 
1. **C  -**  it is also referred as 21st century C

   * inherits expression syntax, basic datatypes, **call-by-value**, parameter passing, pointers
   * most importantly C's emphasis to compile to efficient machine code and communicate naturally with current OS. 


2. **Pascal and decendents -** 

   * modula-2 - inspired package concept from here.
   * Oberon - removed diff b/w module interface file and impl files.
   * Oberon-2 - inspired syntax for packages, imports, declerations( particularly method declerations).


3. **CSP and CSP impls -** originally CSP(Communicating Seq. Processes) was created in bell labs where a program in CSP was described to be parallel composition of processes ***with no shared state***, communicating with each other via channels. But initial CSP was just a formal language to describe the fundamental concept of concurrency.

   * Squeak - first impl of CSP to handle mouse and keyboard events, with statically created channels.
   * NewSqueak - it provided C like expression and Pascal like type notation and here ***channel became first class values, i.e. dynamically created value and storable in variables***

**PS -**  Also has new features like Slices (dynamic arrays with efficient random access but also sofisticated sharing arrangements like Linked Lists) and Defer.


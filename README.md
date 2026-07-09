# GO Lang
- GO is an compiled programming language whose execution speed is comparable to JAVA and C#.
- How to write your first GO program ?
   
## Basics
```
Step 1: Installation. Visit the official page and download based on your local machine.
Step 2: Create a file with ".go" extension.
Step 3: Write code and run
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
- The `package main` is used for defining that this page is what needs to be compiled for execution
- `main` function shows where the entry point is for the execution.
- We need to import `fmt` as its used for printing and scanning from client (fmt stands for format).

*Note*: So when we create a go program then build it and it will create an executable file which can be used anywhere without even insalling GO

## Memory Management
- In languages such as Rust or C we have to do manual memory management.
- In language such as JAVA it has automated memory management.
- In GO there is autmoated memory management but it doesn't have a virtual machine rather a automated runtime. Which means with every bit of binary code has some extra code for memory management.

## Data Types
### Integers
- int8 / uint8: 8-bit size.
- int16 / uint16: 16-bit size.
- int32 / uint32: 32-bit size.
- int64 / uint64: 64-bit size.
- int / uint: Platform-dependent width (32 bits on 32-bit systems, 64 bits on 64-bit systems). `(This is the default choice for whole numbers.)`
*Note*: int have both -ve and +ve where as uint has only +ve numbers.

### String
- string: An immutable sequence of UTF-8 encoded bytes representing text.

### Floating Point 
- float32: 32-bit single precision.
- float64: 64-bit double precision. `(This is the default float type.)`
  
### Special Built-in Aliases
- byte: Alias for uint8. Used for raw binary data (Like JSON or file read data).
- rune: Alias for int32. Used to represent a single Unicode code point.
- uintptr: Unsigned integer large enough to store a pointer address.

## Defining and Assigning a variable
```
var variable_name dataType = value
```
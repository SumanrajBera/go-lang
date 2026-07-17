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

## Type inference
- We make use of `:=` to infere type based on the value
- Example: number := 1 - This is int and username := "" - This is a string. The default inference for float is float64.
- Though its inferred it can't be changed as its a static type language.
- With `fmt.Printf` we can make use of **%T** to get the type of the variable.
- We can convert by using specific types with brackets for type conversion. Example: float64(num) - num is int

## Constant variables
- We use the `const` keyword for making constant variables.

*Note*: = vs :=. Use `=` at package level and use `:=` can be used only inside functions for type inference.

## Interpolation
- Here we can make use of `Printf` or `Sprintf` where Printf is used for printing formatted string to console and Sprintf will return the formatted string
- Types used are as follows:
```
%v - For any value can be used
%d or variations - For int based values
%f or variations (%.f[for rounding], %.nf [For how many decimals]) - For float64 and float32
```

## Comparison operators
- == : For comparison
- != : For not equal to
- ">" : Greater than
- "<" : Smaller than
- ">=" : Greater than equal to
- "<=" : Smaller than  equal to

## Unique Conditional
```
if initialize; compare {condition} - When varible is not used other than here
```
*Note*: We can pass functions in another functions just like callbacks we use in JS.
*Note*: In Go, variables are passed value and not by reference.
*Note*: In Go, it doesn't allow us to have unused variables

## Ignoring return values
- This is important sometimes because compiler can throw error due to unused variables 
```
func getPoint() (x int, y int) {
    return 3, 4
}
x, _ := getPoint() // Here we are telling we know there's a value but we ignore it
```

*Note*: We can return named values as it helps with documentation and improving readability.

## Naked return (Also called implicit return)
```
func getPoint() (x, y int) {
    // Here when we return we will get x and y values this is called naked return
	return
}
```

*Note*: Guard clauses is an early return from the function when a given condition is met.

## Structs 
- A struct is used to define structured data and is used for grouping variables pf different data types togeather.
```
type struct_name struct {
    // Group of variables
}
```
- Nested Structs are structs within structs
```
type sendMessage struct {
	message  string
	sender   user
	receiver user
}

type user struct {
	id   int
	name string
}
```
- Anonymous structs and these are structs which we declare directly or declare and assign directly
```
// Direct declaration
type car struct {
    Wheel struct{
        pressure float32
        company string
    }
}

// Declare and assign
myCar := struct {
    name string
    engine string
    bags int
} {
    name: "Mercedes",
    engine: "New"
    bags: 4
}
```
- Embedded struct is the struct that allows us to make another struct become part of the struct we have embedded it in.
```
type vehicle struct {
    model string
    name string
}

type truck struct {
    // vehicle is embedded so to access model we just need truck.model and not truck.vehicle.model
    vehicle
    tyres int
}
```

## Interface
- Interfaces in GO are abstract methods (we don't use keyword abstract rather `name and return type`) and has implicit implementation (we don't use implements keyword).
- We need to apply all the methods defined in the interface or else the code will fail.
```
// No abstract keyword
type message interface {
    sendMessage() string
}

type report struct {
    id int
    message string
}

func (r report) sendMessage() string {
    return r.message
}

func main() {
    newReport := report {
        id: 123455,
        message: "Hello World",
    }
    fmt.Println(newReport.sendMessage())
}

```
- Empty interface `interface{}` is implemented by every type and we can't use it. These are use to assign values on the go like incoming data from an API which we might not know.
- We can also specify named arguments in the interface method's arguments.
- Type assertion can be used to check if the specified type is part of the interface or not.
```
var i interface{} = "Hello"

s, ok := i.(string) // This will check if interface has type string and if it has then true will be assigned to ok else false
```
*Note*: Keep interfaces small. Interfaces should not be aware of the types they satisfy (I have a vehicle interface and we have a method called isCar `This shouldn't happen because we may then define isTruck, isBike etc.` rather create a sub-interface with the vehicle interface)

## Errors
- We don't have try and catch rather so we don't throw error rather we return and check.
```
import (
    "fmt"
    "errors"
    )

func checkPassword(password string) (bool, error) {
    if len(password) < 8 {
        return false, errors.New("Password can't be smaller than 8 chars")
    }

    return true, nil
}

func main() {
    correct, error := checkPassword("1234567")
    if error != nil {
        fmt.Println("Error", error)
        return 
    }
    
    if !correct {
        fmt.Println("Password isn't correct")
        return
    }

    fmt.Println("Password is correct")
}
```
- We also have `Errorf` is the function that is used for formatting the error comes with `fmt` package.
- We have the error interface and we can use it to create custom errors.
```
// Error interface built-in
type error interface {
    Error() string
}
```
```
// Custom error
import (
    "fmt"
    "errors"
)
type userError struct {
    name string
}

func (userErr userError) Error() string {
    return fmt.Sprintf("%v username exists choose another", userErr.name)
}
```

## Loops
- for loop - It is similar to for loop we see in other languages but it doesn't use parenthesis just like if-else
```
for INTIAL; CONDITION; AFTER {
    // code 
}
```
- There is **no** `while loop` in go and if we use our for loop with only condition it becomes a while loop.
- There is `break` and `continue` statement for loop skiping and quiting.

## Logical AND and OR
- Logical AND is represented by `&&`
- Logical OR is represented by `||`
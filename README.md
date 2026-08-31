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

## Arrays
- It is a fixed size collection of similar data types.
```
var array [3]int = [3]int{1,2,3}
array := [3]int{1,2,3}
array := [...]int{1,2,3}
```
- It is different from JS arrays as its not flexible and can only store upto defined size.

## Slices
- Its still rigid in terms of type but flexible on the side as we can store more than whatever limit we define. We can also make a shallow copy stroed using slice syntax.
```
s := []int{1, 2, 3}

// Using make for performance and optimization (Used when we have a rough idea of size but not sure so we need flexibility)
Syntax: make([]type, len , capacity)
len - for number of elements currently in slice
capacity - how much capacity it can hold before 'GO' needs to automatically allocate space
s := make([]int, 3, 5) - we can skip capacity 

// From array
arr := [3]int{1, 2, 3}
slice := arr[1:3] // Shallow copy from array from idx 1 to 2

// For empty slice
var slice []int  // slice is nil but we need to append to add items
```
- **Append** function is the method that can be used to append values in from the front of the array.
```
Syntax: slice = append(slice, value1, value2, ...values)
```
*Best Practice*: Whenever you append, append and assign to the same which can cause to some bugs. For example, if we don't mention capacity to slice and we append then every time and assign to a different slice its ok as its reallocation is happening in a new space but if capacity was mentioned then it will already have pre-allocated space and appending will only cause it to change the value in the same space.
```go
// No cap allocation is different
i := make([]int, 3)
j := append(i, 4) // [0 0 0 4]
k := append(i,5) // [0 0 0 5]

// With cap allocation is 
i := make([]int, 3, 8)
j := append(i, 4) // [0 0 0 4]
k := append(i,5) // [0 0 0 5]
// But if we print j again you will see its changed to [0 0 0 5] because we already have allocated space
```

## Variadic function (rest and spread operator)
- Variadic functions are basically when we pass multiple arguments of smae type without a count.
```go
// Sum of passed numbers
function sum(nums ...float64) float64 {
    func sum(nums ...float64) float64 {
	var res float64 = 0
	for _, num := range nums {
		res += num
	}
	return res
}
}
```
*Note*: In the above example, as you can see I have made use of range for loop which is shorter version to go through all elements in an array.
- By this we can also pass different interfaces to the function.

## Maps
- Maps are another data structures that is common in every language which is used to store elements in key-value pairs
```go
// Syntax
agesMap := make(map[string]int)

ages := map[string]int {
    "henry": 20,
    "matt": 21,
}
```
- Creating, accessing and updating values in a map
```go
// Creating and updating
map["key"] = elem

// Accessing
fmt.Println(map["key"])
```
- Deleting key and value in a map
```go
delete(mao, "key")
```
- Check if a key exists in a map and get a value 
```go
value, ok = map["key"]
// ok is boolean which is true if it exists and false if it doesn't
```
*Note*: Using Struct is better than making use of nested Maps which are very complicated. As struct is something we can use as the keys. Known schema (users, products, orders, etc.) → Use structs. Unknown or dynamic schema → Use nested maps. 
- Composite type keys with struct 
```go
type ScoreCard struct {
    Place, Team string
}

matches := make(map[ScoreCard]int)

// Keys are matched with each value passed in struct
matches[ScoreCard{"Oklahoma", "Austin"}]++
matches[ScoreCard{"Oklahoma", "Dallas"}]++
matches[ScoreCard{"NYC", "Yorkers"}]++
```
*Note*: If a key doesn't exist and we try to access it then we get 0.

### Nested Maps syntax
```go
nestedMap := make(map[string]map[string]int)

// Without this it will panic (Its a runtime error)
if nestedMap["Oklahoma"] == nil {
	nestedMap["Oklahoma"] = make(map[string]int)
}

// For accessing and updation
nestedMap["Oklahoma"]["Austin"]++
```

## Advanced Functions
- Functions are **first-class citizens** meaning they can be used as variables.
- **Higher-order functions** are the functions that return a function or take function as an argument.
- **Function currying** is a technique in which we take function as an input and return a new function as output.
- *defer* keyword is used to execute a function at the end of current executing function. Mostly used for cleanups.
- **Closures** are used for preserving states. They work in a way that when there is an inner which can access the variables of an outer function. 

## Pointers
- `&` (address-of operator) returns the memory address of a variable.
- `*` (dereference operator) accesses or modifies the value stored at the address held by a pointer.

```go
x := 5

p := &x // p stores the address of x

fmt.Println(p)  // Address stored in p
fmt.Println(*p) // Value stored at that address (5)

*p = 10         // Modify the value at that address
fmt.Println(x)  // 10
```
*Note*: If a pointer points to nothing then its zero value or say default value is **nil**

## Packages
- A package is simply a collection of related Go files that work together.
*Note*: We cannot have different packages in same folder
- The package name must be the **directory name** because that is how we import and in `go.mod` we need to mention the module we want
```
D:.
│   arithmetic.go
│   go.mod
│   
└───packages
        sum.go
For the above structure we need in go.mod file we must have: module packageInGO
So we can import: import "packageInGO/packages"
```
**Note:** `package main` is a special package in Go. It marks an executable program and must contain a `main()` function, which is the program's entry point. Unlike most packages, its name is always `main`, regardless of the directory name.

*Note*: So function which we want to export needs to have uppercase letter at start. GO doesn't have a proper export keyword.

### Modules
- **Modules**: A module is a collection of related packages that are versioned together.
- **Module Location**: There can be only one module per project directory tree, and its go.mod file must reside strictly in the root directory.
- **Module Initialization**: To initialize a directory as a module, use the command go mod init <module-path> (where <module-path> is the import path, such as ://github.com or a simple name like myproject).
*Note*: Modules make packages importable across projects. So if you want to export a package external you need to make sure it has a `go.mod` file

### Commands for package and modules
- `go mod tidy` - This command is used for packages that have been imported but not mentioned in the mod file. 
- `replace <module_name> => <local_path>` - This command is used for development for files that aren't pushed to github and use them instead.
- `go mod init <module_name>` - This command is used for initialising module
- `got get <module_name>` - To automatically add our packages so that they can be used.
- `cat go.mod` - This command can be used to see the contents of the `go.mod` file.

*Note*: `go.sum` is the alternative to package-lock.json which is used for making sure correct version dependencies are used.

## Channels and concurrency
- **Synchronous code**: This code executes in order from top to bottom line-by-line. Its not the most efficient for writing optimised code.
- **Concurrent code**: Instead of executing one task completely before moving to the next, we can have multiple tasks in progress at the same time. When one task is waiting for something, such as a network or database response, another task can make progress. This allows us to reduce unnecessary waiting and use our resources more efficiently.
  
### Routines
- So routines in go are specified threads which allows our task to run concurrently and is specified using keyword `go`
- Example
```go

func emailingWithGroup(sender string, receiver string, wg *sync.WaitGroup) {
	wg.Add(1)
    // This occurs in different routine then the current
	go func() {
		defer wg.Done()
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("Email received from %s \n", receiver)
	}()

	fmt.Printf("Email sent from %s \n", sender)
}
```
- To make sure our other routines are completed befoe main completes its execution we can use a method from `sync` package which is `WaitGroup` which is a counter and has three things (Add, Done and Wait) which allows main to continue to stay alive until the routines complete their execution. It is passed using pointer.
```go
var wg sync.WaitGroup
// This adds 1 to the counter which can indicate a single routine
wg.Add(1)
// This tells to remove that 1 which means the current routine completed
wg.Done()
// This allows our main function to prolong so that until counter becomes it doesn't complete its execution
wg.Wait()
```

### Channels
- Channels allow goroutines to communicate values/results that create dependencies between pieces of work.
- It is a type which can be created `chan` keyword and the data that we want to pass must be sent using `<-` operator which allows the data to flow in the direction we are pointing.
```go
processOrder := make(chan Order) // Channel with Order as type
processOrder <- order // Passing order into channel (Here we can create a go routine which allows to process values passed to channel)
```
*Note*: An **unbuffered** channel's (no storage) requires the sender and receiver to synchronize, so a send waits until someone receives it. A **buffered** channel has a fixed capacity, so sent values can remain in the buffer until a receiver consumes them, acting like a bounded queue.

- To check if a channel is closed we can use the type assertion technique with simple modification
```go
v, ok := <- ch // ok will return false if the channel is closed  or if its empty (buffered channel)
```
- We can close a channel to indicate that we are done with the channel
```go
close(ch)
```
### Select (swicth statement for channels)
- `select` is the switch statement execlusive to channels.
- We use `select` when a Goroutine needs to wait on multiple asynchronous operations.
```go
for {
	select {
		case val, ok := <- firstChan { 
			if !ok {
				return
    		}
			fmt.Println(val)
        }
		case val, ok := <- secondChan {
			if !ok {
				return
    		}
			fmt.Println(val)
        }
    }
}
```
- We can also use **default** case which allows us to do something if no channel is ready.
- If all channels are ready go uses a random generator to decide which case to execute.
  
### Read-only and Write-only channels
- Read-only and Write-only are only used for reading or writing from and to the channel
```go
// Read-only channel
func readChannel(ch <-chan int) {
	val := <-ch
	fmt.Println("The value in channel:", val)
}

// Write-only channel
func writeChannel(val int, ch chan<- int) {
	ch <- val
}
```
## Mutexes
- A `mutex` stands for Mutual Exclusion which means we are agreeing to mutually exclude something.
- This is used for synchronisation and avoiding data race condition which happens when two routines try to modify same data structure at same time.
- With mutexes we can `lock` and `unlock` the data structure so that other routines can't access it until the mutex's lock isn't unlocked.
```go
var mux sync.Mutex
// To lock a part of code
mux.Lock()
// To unlock a part of code
mux.Unlock()
```

### RWMutex
- This is part of the same package but comes with some extra locks and unlocks for reading and writing as the name suggests.
- The following are the locks and unlocks tha come with **RWMutex**
  - RLock - This is used for locking from writing so that until its released no routines can write in the locked resource. This locked can be shared simulatneously by multiple routines which means until each one is released we can't write.
  - RUnlock - This is used for unlocking RLock.
  - Lock - This is used for locking reading and writing for other routines and once released then only can other routines use the resouces.
  - Unlock - This is used for unlocking Lock.
- We can run into deadlocks if we're not careful. For example, if a goroutine holds an RLock() and then tries to acquire Lock(), Lock() will block waiting for the read lock to be released. However, the goroutine can't reach RUnlock() because it's blocked on Lock(), resulting in a deadlock.
```go
rw.RLock() // Suppose we RLock here

// Then we try to acquire the write lock in the same goroutine
rw.Lock() // Blocks because the RLock must be released first, but this goroutine is stuck here and can't reach RUnlock()

rw.RUnlock() // Never reached
```
*Note*: The `mutex` is just a way to implement synchronisation but if we are not using it we can still run into some data race conditions. Example, if one routine uses mutex and other doesn't and they both are trying to update it we run into data race condition.

## Generics
- Generics lets us write a function, type, or data structure that works with multiple types without duplicating the implementation.
```go
// Here we can use this add function for both int and float
// If we didn't have generics then we would have to create 2 functions for same task
func add[T int | float64](a T, b T) T {
	return a + b
}
```
### Contraints
- Generics let you work with multiple types; constraints define which types are permitted.
```go
// Without constraints
func midElem[T any](arr []T) T {
	var midIdx int = len(arr) / 2
	return arr[midIdx]
}

// With Contraints for accepting int and float64
func add[T int | float64](a T, b T) T {
	return a + b
}
```
### Interface Types List
- We can use [T any] when we want a generic function to accept any type. If we want to restrict T, we can directly define a constraint such as [T ~string | ~int].
- However, if we need to use the same constraint across multiple functions, writing it repeatedly would violate the DRY principle.
- We can instead define a reusable interface as a type constraint and use it across our generic functions.
```go
type Ordered interface { ~string | ~int }

func Min[T Ordered](a, b T) T { // ... }

// This allows us to define the constraint once and reuse it wherever needed.
```
### Parametric Constraints
- We can make use of `interfaces` as a constraint on a type parameter to make sure we only accept types that satisfy (implement) that interface.
```go
type store[P product] interface {
    Sell(P)
}

type product interface {
    Price() float64
    Name() string
}
// Here, P can only be a type that satisfies the product interface.
```
*Note*: The name of a type parameter can be anything but `T` is a common convention.

Golang Cheatsheet
===================================

## Golang links
1. [Download GO](https://golang.org/dl/)
2. [Golag Playground](https://play.golang.org/)
3. [Golag Docs](https://golang.org/pkg/)
4. [Go Dev](https://go.dev/)

## About Golang
* Go was created by Google and released about 10 years ago in order to run on their servers and is focused on code readability.
* Go is compiled language which generates platform-specific binaries.
* Go is statically typed. That means that the compiler knows at compile-time the types of our variables and
* `go` binary is the standard tool of the language and is used to perform various tasks
```
	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages
```
* Use `go fmt` to format/lint your code (or `goimports` which also takes care of your imports)
	1. can perform checks on them to verify our code
	2. can optimize for memory usage
* "Tabs or Spaces?" Go uses tabs to indent your code (this is imposed by `go fmt`)

## The typical "Hello world!"
```
package main

import "fmt"

func main() {
    fmt.Printf("Hello world!")
}
```

## Variables/Constants
### Types
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32 ~= a character (Unicode code point)

float32 float64

complex64 complex128
```

### Operators
### Arithmetic
|Operator|Description|
|--------|-----------|
|`+`|addition|
|`-`|subtraction|
|`*`|multiplication|
|`/`|quotient|
|`%`|remainder|
|`&`|bitwise and|
|`\|`|bitwise or|
|`^`|bitwise xor|
|`&^`|bit clear (and not)|
|`<<`|left shift|
|`>>`|right shift|

### Comparison
|Operator|Description|
|--------|-----------|
|`==`|equal|
|`!=`|not equal|
|`<`|less than|
|`<=`|less than or equal|
|`>`|greater than|
|`>=`|greater than or equal|

### Logical
|Operator|Description|
|--------|-----------|
|`&&`|logical and|
|`\|\|`|logical or|
|`!`|logical not|

### Examples
```
var foo1 int // declaration without initialization

var foo2 int = 42 // declaration with initialization
var foo3 = 42     // type will be inferred
foo4 := 42        // shorthand, only in func bodies
fmt.Printf("%d\n", foo1)
fmt.Printf("%d\n", foo2)
fmt.Printf("%d\n", foo3)
fmt.Printf("%d\n", foo4)

var foo5, bar int = 42, 1302 // declare and init multiple vars at once
fmt.Printf("%d\n", foo5)
fmt.Printf("%d\n", bar)

const constant = "This is a constant"

// iota can be used for incrementing numbers, starting from 0
const (
    _ = iota      // 0 is skipped
    a             // 1
    b             // 2
    c = 1 << iota // 8 (2^3)
    d             // 16 (2^4)
)
fmt.Printf("%d\n", a)
fmt.Printf("%d\n", b)
fmt.Printf("%d\n", c)
fmt.Printf("%d\n", d)

// bool
var b1 bool = true
fmt.Printf("%t\n", b1)

// string
var s1 string = "text"
fmt.Printf("%s\n", s1)

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
var i1 uint16 = 65_535
fmt.Printf("%d\n", i1)

// byte // alias for uint8
var b2 byte = 255
fmt.Printf("%d\n", b2)

// rune // alias for int32 ~= a character (Unicode code point)
var r1 rune = 'a'
fmt.Printf("%v\n", r1)         // 97
fmt.Printf("%s\n", string(r1)) // a

// float32 float64
var f1 float64 = 3.141_592_653_589_793
fmt.Printf("%.15f\n", f1)

// complex64 complex128
var c1 complex64 = 3.2 + 5.4i
fmt.Printf("%v\n", c1) // (3.2+5.4i)
```

### Conversions
#### int - float
```
i := 42
f := float64(i)
u := uint(f)
```

#### int - string
```
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
s := strconv.FormatInt(-42, 10)  // params (int, base)
i, err := strconv.ParseInt("-42", 10, 64) // params (int, base, return size)
s := strconv.FormatUint(42, 16)
u, err := strconv.ParseUint("42", 10, 64)
```

#### bool - string
```
s := strconv.FormatBool(true)
b, err := strconv.ParseBool("true")
```

#### float - string
```
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
f, err := strconv.ParseFloat("3.1415", 64)
```

#### rune - string
```
s := string(rune1)
r := rune(str[2])
```

#### []byte - string
```
b := []byte("some text")
s := string(bytes)
```

#### ascii - unicode
```
q := strconv.Quote("Hello, world!")
q := strconv.QuoteToASCII("Hello, 世界")
uq, err := strconv.Unquote(q)
```

[strconv documentation](https://golang.org/pkg/strconv/#ParseInt)

## Flow control
### Examples - if
```
if var1 := a + b; var1 < 100 {
	fmt.Println(var1)
} else if var1 == 100 {
	fmt.Println(0)
} else {
	fmt.Println(var1 - 100)
}
```

### Examples - for
#### Typical loops
```
for i := 1; i < 10; i++ { // for loop
	fmt.Println("test")
}
```

```
for i < 10  { // while loop
}
```

```
for { // infinite loop
}
```

#### Breaking out of loops
```
for i := 5; i < 20; i++ {
	for j := 0; j < i; j++ {
		fmt.Println("")
		if j > 15 {
			break
		}
		fmt.Print(i, j)
		if j > 10 {
			continue
		}
		fmt.Print("-")
	}
}

println("\n\n\n")
```

```
mark1:
	for i := 5; i < 20; i++ {
		for j := 0; j < i; j++ {
			fmt.Println("")
			if j > 15 {
				continue mark1
			}
			fmt.Print(i, j)
			if j > 10 {
				break mark1
			}
			fmt.Print("-")
		}
	}
```

### Examples - switch
```
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("Mac OS Hipster")
case "linux":
	fmt.Println("Linux Geek")
default:
	fmt.Println("Other")
}

number := 7
switch {
	case number < 7:
		fmt.Println("smaller")
	case number == 7:
		fmt.Println("equal")
	case number > 7:
		fmt.Println("greater")
}

byte := 'a'
switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("a vowel")
	default:
		fmt.Println("a consonant")
}
```

### Examples - +1 special for concurrency
*UPCOMING*

## Functions
### Examples - Basics
```
func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func plusMinus(a int, b int) (int, int) {

	return a + b, a - b
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res1, res2 := plusMinus(3, 2)
	fmt.Println("3+2 =", res1)
	fmt.Println("3-2 =", res2)
}
```

### Examples - Nameless functions
```
func main() {
    // assign a function to a name
    add := func(a, b int) int {
        return a + b
    }
    // use the name to call the function
    fmt.Println(add(3, 4))
}
```

### Examples - Closures
```
func outer() (func() int, int) {
    outer_var := 2
    inner := func() int {
        outer_var += 99 // outer_var from outer scope is mutated.
        return outer_var
    }
    inner()
    return inner, outer_var // return inner func and mutated
}

func main() {
    // use the name to call the function
    add99, innerValue := outer()
    fmt.Printf("%d %d", add99(), innerValue)
}
```

### Examples - Variadic
```
func adder(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(adder(1, 2, 3)) 	// 6
	fmt.Println(adder(9, 9))	// 18

	nums := []int{10, 20, 30}
	fmt.Println(adder(nums...))	// 60
}
```

### Examples - Defer
```
func closer(suffix string) {
	fmt.Println("Closing Step "+suffix)
}

func main() {
	fmt.Println("Step 1")
	defer closer("1")
	fmt.Println("Step 2")
	defer closer("2")
}
```
*will print*
```
Step 1
Step 2
Closing Step 2
Closing Step 1
```

## Arrays
### Examples - Arrays
```
var a [5]int // declare an int array with length 5
a[0] = 42 // arrays are zero-based
i := a[0]

// declare and initialize
a := [5]string{"a", "b", "c", "d", "e"}
a := [...]int{"a", "b", "c", "d", "e"} // let compiler figure out array length
```

## Slices
Slices are references to an underling arrays, giving the impression of an
array. They provide a way to work with sequencies of typed data.

The zero value of a slice is `nil`.

### Memory structure
In the background slices have:
* a pointer to the underling array
* a length counter which can be retrieved by `len(slice)`
* a capacity counter which can be retrieved by `cap(slice)`

### Examples - Creating Slices
```
slice1 := []string{"a", "b", "c"}

// create an underling [10]int array and set the length to 5
slice2 := make([]int, 5, 10)
fmt.Println(len(slice2))  // 5
fmt.Println(cap(slice2))  // 10

array1 := [3]string{"one", "two", "three"}
slice3 := array1[:] // create a slice on array1
slice4 := array1[1:3] // create a slice on items 0,1 of array1
```

### Examples - Iterate over Slices
```
for index, value := range slice {
	fmt.Println(index, value)
}

// Iterate over values
for _, value := range slice {
	fmt.Println(index, value)
}
```

### Examples - Append to Slices
```
slice1 := []string{"one", "two", "three"}
slice2 := []string{"1", "2", "3"}

slice1 = append(slice1, "four") // append single value
slice1 = append(slice1, slice2...) // append an other slice
```

### Examples - Copy Slices
```
slice1 := []string{"one", "two", "three"}
slice2 := []string{"1", "2", "3", "4"}

c := copy(slice2, slice1)  // slice1 == []string{"1", "2", "3"},  c == 3
```


## Maps
Zero value of map is `nil`

### Examples - Create maps
```
map1 := map[string]bool{
	"key1": true,
	"key2": true,
}

map2 := make(map[string]boolean) // Create an empty map. i.e. len(map2) == 0
```

### Examples - Editing maps
```
map1 := map[string]bool{
	"key1": true,
	"key2": true,
}

map1["key3"] = false  // create and set a new key:value pair

delete(map1, "key2")  // delete key2
```


## Structs
### Examples
```
// Declaration
type Vertex struct {
    X, Y int
}

// Creating
var v = Vertex{1, 2}
var v = Vertex{X: 1, Y: 2} // Creates a struct by defining values with keys
var v = []Vertex{{1,2},{5,2},{5,5}} // Initialize a slice of structs

// Accessing members
v.X = 4
```

## Packages
* Go will treat every directory under the project root which contains .go files as a package
* All .go files in a package are concatinated(smart concatination) before compilation
* Package members (e.g. vars, funcs, ...) MUST start with capital letter
* When importing package we must always use the full package names

### Examples - imports
```
// Singleline import. (Compiler will accept this but "go fmt" will dissagry)
import "fmt"
import "log"

// Multiline import
import (
    "fmt"
    "log"
    "os"
)

// Multiline import, with alias
import (
    "fmt"
    standarLog "log"
    "os"
)
```

## Pointers in go
* Avoid using pointers in any other type of variable than *struct*
* Use `&` in front of a variable name to get a pointer to the variable
* Use `*` in front of a pointer variable to get the value of this pointer
* Use `*<TYPE>` as a variable type to designate a pointer to this type of values

### Example
```
func main() {
	type Coord struct {
		X float32
		Y float32
	}

	c1 := Coord{X: 1, Y: 2}
	c2 := &c1              // <<< Get the memory address to c1 and put it in pc1
	fmt.Printf("%T\n", c2) // *main.Coord

	fmt.Println(c1) // {1 2}
	fmt.Println(c2) // {1 2}
	c1.X = 5
	fmt.Println(c1) // {5 2}
	fmt.Println(c2) // {5 2}

	add := func(a, b Coord) Coord {
		return Coord{a.X + b.X, a.Y + b.Y}
	}
	increase := func(a *Coord, b Coord) { // <<< Use pointer as a parameter
		a.X += b.X
		a.Y += b.Y
	}

	c3 := Coord{10, 20}
	res1 := add(c1, c3)
	c1.X++
	fmt.Println(c1)
	fmt.Println(res1)

	increase(&c1, c3) // <<< Pass a pointer parameter
	fmt.Println(c1)

	pc3 := &c3
	increase(&c1, *pc3) // <<< Pass a pointer as a value parameter (pc3)
	fmt.Println(c1)
}
```

## Methods
* `Pointer semanctics` when we are using pointers as return type to the New... function and as the receiver to our methods
* `Value semantics` when we are using values as return type to the New... function and as the receiver to our methods
* We should avoid using mixed semantics when we are implementing our methods for better readability. (but, some times it is required and we cannot avoid it)

```
type Coord struct {
	X float32
	Y float32
}

func NewCoord(x, y float32) Coord {
	return Coord{x, y}
}

func (c Coord) add(b Coord) Coord { // <<< Use value as a receiver
	return Coord{c.X + b.X, c.Y + b.Y}
}
func (c *Coord) increase(b Coord) { // <<< Use pointer as a receiver
	c.X += b.X
	c.Y += b.Y
}

func main() {
	c1 := NewCoord(1, 2)
	c2 := &c1
	fmt.Printf("%T\n", c2) // *main.Coord

	fmt.Println(c1) // {1 2}
	fmt.Println(c2) // {1 2}
	c1.X = 5
	fmt.Println(c1) // {5 2}
	fmt.Println(c2) // {5 2}

	c3 := NewCoord(10, 20)
	res1 := c1.add(c3)
	c1.X++
	fmt.Println(c1)
	fmt.Println(res1)

	c1.increase(c3)
	fmt.Println(c1)

	pc3 := &c3
	c1.increase(*pc3)
	fmt.Println(c1)
}
```


### Creating a package (for Go >=1.11.1)
Type the following command in your application's route if `go.mod` does not exit
```
go mod init APPLICATION_NAME
```

Remember to add the `package PACKAGE_NAME` in all .go files.

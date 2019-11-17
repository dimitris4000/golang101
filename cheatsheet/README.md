Golang Cheatsheet
===================================

## Golang links
1. [Download GO](https://golang.org/dl/)
2. [Golag Playground](https://play.golang.org/)
3. [Golag Docs](https://golang.org/pkg/)
4. [Go Dev](https://go.dev/)

## Variables/Constants
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

## Flow control
### Examples - if
### Examples - for
### Examples - switch

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


## Packages
### Examples

## Arrays, Slices
### Examples

## Maps
### Examples

## Structs
### Examples

Golang Cheatsheet
===================================
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
```
// int - float
i := 42
f := float64(i)
u := uint(f)

// int - string
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
s := strconv.FormatInt(-42, 10)  // params (int, base)
i, err := strconv.ParseInt("-42", 10, 64) // params (int, base, return size)
s := strconv.FormatUint(42, 16)
u, err := strconv.ParseUint("42", 10, 64)

// bool - string
s := strconv.FormatBool(true)
b, err := strconv.ParseBool("true")

// float - string
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
f, err := strconv.ParseFloat("3.1415", 64)

// rune - string
s := string(rune1)
r := rune(str[2])

// []byte - string
b := []byte("some text")
s := string(bytes)

// ascii - unicode
q := strconv.Quote("Hello, world!")
q := strconv.QuoteToASCII("Hello, 世界")
uq, err := strconv.Unquote(q)

[strconv documentation](https://golang.org/pkg/strconv/#ParseInt)
```

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

## Functions
### Examples

## Flow control
### Examples

## Packages
### Examples

## Arrays, Slices
### Examples

## Maps
### Examples

## Structs
### Examples

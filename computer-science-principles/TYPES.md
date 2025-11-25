# Go Primitive Types

In Go, primitive types are the basic data types that are not composed of other types. They are the building blocks for more complex data structures. These types are also sometimes referred to as "basic types".

## Booleans

Boolean types are used to represent truth values.

- `bool`: Represents a boolean value, which can be either `true` or `false`.

**Usage:** Used for conditional logic and flags.

**Example:**
```go
var isStudent bool = true
```

## Numerics

Numeric types are used to represent numbers. They are further divided into integers, floating-point numbers, and complex numbers.

### Integers

Integers are whole numbers. Go provides both signed and unsigned integers of various sizes.

- `int8`, `int16`, `int32`, `int64`: Signed integers of 8, 16, 32, and 64 bits respectively.
- `uint8`, `uint16`, `uint32`, `uint64`: Unsigned integers of 8, 16, 32, and 64 bits respectively.
- `int`: The size is platform-dependent (32 or 64 bits). It is the most commonly used integer type.
- `uint`: The size is platform-dependent (32 or 64 bits).
- `rune`: An alias for `int32`. It is used to represent a Unicode code point. While a `byte` represents a single byte, a `rune` can represent a multi-byte character.

- `byte`: An alias for `uint8`. It is used to represent a single byte.
- `uintptr`: An unsigned integer type large enough to store the uninterpreted bits of a pointer value. It is used in advanced and low-level programming, often with the `unsafe` package, for operations on memory addresses.

**Usage:** Used for counting, indexing, and calculations involving whole numbers.

**Example:**
```go
var age int = 30
var temperature int32 = -5
```

**`rune` Example:**
When you are working with strings that contain multi-byte characters (like emojis or characters from non-English languages), you should use `rune` to correctly handle each character.

```go
import "fmt"

func main() {
    str := "Hello, 世界"
    for _, r := range str {
        fmt.Printf("%c ", r)
    }
}
// Output: H e l l o ,   世 界
```

**`uintptr` Example:**
`uintptr` is typically used for low-level memory operations and is not something you will use in everyday Go programming. It's often used with the `unsafe` package, which, as the name suggests, can be dangerous if not used correctly.

```go
import (
    "fmt"
    "unsafe"
)

func main() {
    var x int = 10
    // Get the memory address of x
    p := &x
    // Convert the pointer to a uintptr
    uintp := uintptr(unsafe.Pointer(p))
    fmt.Println("Memory address of x:", uintp)
}
```

### Floating-Point Numbers

Floating-point numbers are used to represent decimal numbers.

- `float32`: A 32-bit floating-point number.
- `float64`: A 64-bit floating-point number. `float64` is the default floating-point type and is generally preferred for accuracy.

**Usage:** Used for calculations involving fractions and decimal values.

**Example:**
```go
var price float64 = 29.99
```

### `float32` vs `float64` (Float vs Double)

In many programming languages, you'll find `float` and `double` types for floating-point numbers.

- `float` is typically a single-precision (32-bit) floating-point number.
- `double` is a double-precision (64-bit) floating-point number.

In Go:
- `float32` is the equivalent of a `float`.
- `float64` is the equivalent of a `double`.

The key difference is precision. `float64` can represent a much larger range of numbers and has more precision (more digits after the decimal point) than `float32`. Because of this, `float64` is the default and recommended type for most floating-point calculations in Go to avoid precision errors.

### Complex Numbers

Complex numbers are numbers that have two parts: a **real** part and an **imaginary** part. They are written in the form `a + bi`, where:
- `a` is the real part.
- `b` is the imaginary part.
- `i` is the imaginary unit, which is defined as the square root of -1 (`√-1`).

In Go, complex numbers are represented by the `complex64` and `complex128` types.

- `complex64`: A complex number with `float32` real and imaginary parts.
- `complex128`: A complex number with `float64` real and imaginary parts.

**Usage:** Used in scientific, engineering, and mathematical applications, especially in fields like electrical engineering and signal processing.

**Example:**
You can create a complex number and access its real and imaginary parts using the `real()` and `imag()` built-in functions.
```go
import "fmt"

func main() {
    var c complex128 = 2 + 3i
    
    // Get the real part
    realPart := real(c)
    fmt.Println("Real part:", realPart) // Output: Real part: 2
    
    // Get the imaginary part
    imaginaryPart := imag(c)
    fmt.Println("Imaginary part:", imaginaryPart) // Output: Imaginary part: 3
}
```

## Strings

Strings are used to represent a sequence of characters.

- `string`: Represents a sequence of characters. Strings in Go are immutable, meaning they cannot be changed once created.

**Usage:** Used for storing and manipulating text.

**Example:**
```go
var name string = "Alice"
```

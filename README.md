# Go Basics Guide 🏆

Practical Go programming guide covering the basics and some advanced topics. This guide is designed to help you quickly get started with Go and understand its core concepts.

---

### 1. Hello World

- Print "Hello World!"

---

### 2. Variables

- Declare and Initialize Variables
- Data Types
- Constants
- Printing Variables
- `iota` and Enumerated Constants

---

### 3. Basic Functions

- Declare Functions
- Function Parameters and Return Values
- Naked Returns
- Use Functions Across Files (Same Package)

---

### 4. Conditionals

- `if` / `else` Statements
- `switch` Statements
- `switch true` Pattern (Conditional Cases)

---

### 5. Operators

- Arithmetic Operators
- Comparison Operators
- Logical Operators
- Bitwise Operators
- Assignment Operators
- Increment/Decrement Operators

---

### 6. Loops

- Basic `for` Loops
- Range Loops (`for range`)
- Simulate `while` Loops
- Infinite Loops
- `break` and `continue`

---

### 7. String Manipulation

- Common String Functions (`strings` package)
- String Formatting
- Type Conversion (String ↔ Number) (`strconv` package)
- Regular Expressions (`regexp`)
- Iterating Over Bytes and Runes
- Byte Slice Manipulation (`bytes` package)

---

### 8. Arrays and Slices

- Arrays and Fixed-Length Arrays
- Size Inference with `[...]T`
- Slices and Slice Literals
- Iteration with `range`
- Copying and Appending Slices

---

### 9. Maps

- Declare and Initialize Maps
- Map Literals
- Access and Modify Keys
- Iterate Over Maps
- Delete Keys from Map (`delete()` function)
- Comma-ok Idiom

---

### 10. Advanced Functions

- Function Values (First-Class Functions)
- Assign Functions to Variables
- Pass Functions as Parameters
- Return Functions from Functions
- Anonymous Functions
- Closures
- Variadic Functions (`...T`)
- Recursive Functions

---

### 11. Pointers

- Declare and Use Pointers (`*T`)
- Get the Address of a Variable (`&`)
- Dereference a Pointer (`*p`)
- Modify Values Through Pointers
- Handle Nil Pointers Safely
- Allocate Memory with `new()`

---

### 12. Structs

- Declare Struct Types
- Access and Modify Struct Fields
- Use Pointers to Structs
- Anonymous Structs
- Nested Structs
- Struct Methods with Receivers

---

### 13. Interfaces

- Declare Interfaces
- Implicit Interface Implementation
- Empty Interface (`interface{}`)
- Type Assertions, Type Switches, and Comma-ok Idiom
- Type Aliases
- Built-in Interfaces: `error`, `Stringer`, etc.
- "Explicit" Interface Implementation

---

### 14. Packages

- Create Custom Packages
- Import Packages
- Package Visibility and Capitalization Rules

---

### 15. Error Handling

- Return and Handle `error` Values
- Must Idiom (Short-Circuit Failures)
- Use `panic`, `recover`, and `defer`
- Define Custom Error Types
- Wrap Errors with `%w` and `errors.New`
- Use `errors.Is` and `errors.As`

---

### 16. Generics

- Generic Function Syntax
- Define Generic Struct Types
- Type Constraints: `any`, `comparable`, `~T`
- Union Types and Type Sets
- Generics with Interfaces
- Type Aliases with Generics

---

### 17. Concurrency

- Goroutines (`go` Keyword)
- Channels (Buffered and Unbuffered, Comma-ok Idiom, Direction Types, Iteration)
- `select` Statement
- `sync` Package
  - WaitGroups
  - Mutual Exclusion (Mutexes, RWMutexes - Read/Write Locks)
  - Once
  - Pool
  - Cond
  - Atomic Operations (`sync/atomic` Package)
- `context` Package (Timeouts, Cancellation, Passing Values)
- Deadlocks, Race Conditions (`-race` flag), and Goroutine Leaks
- Concurrency Patterns
  - Fan-out, Fan-in
  - Generator
  - Pipeline
  - Worker Pool
  - Semaphore

---

### TODO:

19. HTTP Servers
    - JSON and XML Handling
    - Data Fetching
    - HTTP Server
    - TCP Server
20. Testing
21. Cryptography

- Environment Variables
- OuterLoop label break
- init() function
- C
- Assembly
- gRPC
- HTMX
- Reflection

## 🧩 General Data Type Categories in Go

In Go, data types are divided into several main categories:

---

### **1. Basic (or Primitive) Types**

Types that directly hold values.
Examples: numbers, strings, booleans.

🧠 Examples:

```go
bool
string
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
complex64, complex128
byte (alias for uint8)
rune (alias for int32)
```

---

### **2. Composite Types**

Types that are composed of multiple simple data elements.

🧠 Examples:

```go
array   // Fixed-size array
slice   // Dynamic array
map     // Key-value pairs
struct  // Data structure containing multiple fields
```

---

### **3. Reference Types**

Instead of holding the value directly, they reference a memory address.

🧠 Examples:

```go
pointer  // Pointer to a value
function // Function as a value
channel  // For communication between goroutines
interface // Abstract type for multiple different types
```

---

### **4. Interface Types**

Types that only define **behavior (methods)**, not data.
Any type that implements those methods belongs to that interface.

🧠 Example:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

---

### **5. Function Type**

In Go, functions themselves are a data type, meaning you can assign functions to variables or pass them as parameters.

🧠 Example:

```go
var f func(int) int
```

---

### 📘 Summary Table:

| Category   | Examples                              | Description              |
| ---------- | ------------------------------------- | ------------------------ |
| Basic      | int, bool, string                     | Fundamental types        |
| Composite  | array, slice, map, struct             | Compound types           |
| Reference  | pointer, function, channel, interface | Reference types          |
| Interface  | Writer, Reader                        | Behavioral types         |
| Function   | func()                                | Function as a type       |

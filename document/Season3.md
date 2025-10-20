### **۱. Methods of Variable Declaration in Go and Their Differences and Uses**

In Go, there are several ways to declare a variable 👇

---

#### 🔹 **Method 1: Using `var`**

```go
var x int = 10
```

- The type is specified.
- You can provide an initial value or not.
- If no value is provided, it gets the zero value.

🧠 Example:

```go
var name string     // Value is "" (empty string)
var age int = 25
```

✅ **Use Case:** When you want to specify the data type or declare a variable at the package level.

---

#### 🔹 **Method 2: Using `var` without Type (Type Inference)**

```go
var x = 10
```

Go automatically infers the variable type (here `int`).

✅ **Use Case:** When the value is clear and you don't want to manually write the type.

---

#### 🔹 **Method 3: Using `:=` (Shorter and More Common)**

```go
x := 10
```

- Can only be used inside a function (not at the package level).
- Go infers the type from the value.
- The most common method for declaring local variables.

✅ **Use Case:** When you are inside a function and want to quickly declare a variable.

---

#### 🔹 **Method 4: Multiple Declaration**

```go
var a, b, c = 1, 2, 3
```

Or

```go
a, b := 5, "test"
```

✅ **Use Case:** When you want to declare or initialize multiple variables at once.

---

#### 🔹 **Method 5: Block Declaration**

```go
var (
  name = "Mohsen"
  age  = 26
  ok   bool
)
```

✅ **Use Case:** To make the code cleaner and declare several related variables together.

---

🧠 **Summary of Differences:**

| Method             | Type Specification | Scope of Use     | Requires Initial Value | Key Feature               |
| ------------------ | ------------------ | ---------------- | --------------------- | ------------------------- |
| `var x int = 10`   | Manual with type   | Everywhere       | Optional              | More control              |
| `var x = 10`       | Type inference     | Everywhere       | Mandatory             | Automatic and readable    |
| `x := 10`          | Short              | Inside functions only | Mandatory        | Fast and common           |
| `var (...)`        | Block              | Everywhere       | Optional              | Clean and organized       |
| `a, b := 1, 2`     | Multiple           | Inside functions only | Mandatory        | Simultaneous declaration  |

---

### **۲. What is `const` in Go?**

`const` is used to define **constants (fixed at compile time)**.
After definition, you cannot change its value.

🧠 Example:

```go
const Pi = 3.14
const Greeting = "Hello"
```

Or multiple:

```go
const (
  A = 1
  B = 2
  C = 3
)
```

📌 Notes:

- You can only assign **constant values (literals)**, not function outputs.
- You can use `iota` to create simple enums.

🧠 Example with `iota`:

```go
const (
  Sunday = iota
  Monday
  Tuesday
)
```

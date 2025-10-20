### 🧩 **۱. Variadic Functions (توابع با ورودی متغیر)**

✅ **Definition:**
Functions that can take **an unspecified number of arguments**.
You don't need to know exactly how many inputs you have.

🧠 **How to define:**

```go
func sum(nums ...int) int {
  total := 0
  for _, n := range nums {
    total += n
  }
  return total
}
```

📦 Calling:

```go
sum(1, 2, 3)
sum(10, 20)
```

✅ **Use Case:**
When a function needs to take many inputs of the same type (like `fmt.Println()` or `sum()`).

---

### 🧠 **۲. Anonymous Functions (توابع بی‌نام)**

✅ **Definition:**
Functions that **have no name** and are usually defined and used **inline (on the spot)**.

🧠 **Example:**

```go
func() {
  fmt.Println("Hello")
}()
```

Or in a variable:

```go
greet := func(name string) {
  fmt.Println("Hi", name)
}
greet("Mohsen")
```

✅ **Use Case:**

- For temporary functions
- For callbacks
- For working with goroutines or event handlers

🧩 Example:

```go
go func() {
  fmt.Println("Running in goroutine")
}()
```

---

### 🧠 **۳. Closure Functions (توابع بسته / دربسته)**

✅ **Definition:**
A function that has access to **variables outside of itself** even after that scope has finished executing.

🧠 **Example:**

```go
func counter() func() int {
  count := 0
  return func() int {
    count++
    return count
  }
}

next := counter()
fmt.Println(next()) // 1
fmt.Println(next()) // 2
```

✅ **Use Case:**

- Creating stateful functions
- Preserving data between multiple executions
- Implementing data encapsulation

🧩 Summary:

> Closure = A function that "remembers" the variables from its outer environment.

---

### 🧠 **۴. Defer Functions (توابع با اجرای تأخیری)**

✅ **Definition:**
The `defer` keyword causes **the execution of a function to be delayed until the current function finishes**.

🧠 **Example:**

```go
func main() {
  defer fmt.Println("World")
  fmt.Println("Hello")
}
```

📦 Output:

```
Hello
World
```

✅ **Use Case:**

- For **resource cleanup** like closing files or connections
- Executing tasks that must be done *at the end* (like finally in other languages)

🧠 Practical Example:

```go
file, _ := os.Open("data.txt")
defer file.Close() // File closes at the end
```

---

### 📘 **Quick Summary Table**

| Function Type | Definition                       | Main Use Case         | Example                      |
| ------------- | -------------------------------- | --------------------- | ---------------------------- |
| **Variadic**  | Input with unspecified count     | Sum or print multiple values | `fmt.Println(a, b, c...)` |
| **Anonymous** | Function without a name          | Callbacks, goroutines | `go func() { ... }()`        |
| **Closure**   | Maintains state from outer scope | Counters, internal memory | `counter()` function         |
| **Defer**     | Executes function at end of current function | Resource cleanup | `defer file.Close()`         |

### **1. What is a Package?**

A package is a collection of Go files that are built together to perform a specific task.
Every Go file starts with `package packagename`.

🧠 Example:
`fmt`, `strings`, and `net/http` are standard Go packages.

---

### **2. What is a Module?**

A module is a higher-level unit than a package; it's a collection of multiple packages that together form a complete project.
Each module is created using the `go mod init` command and has a `go.mod` file.

🧠 Example:
A web project can have several packages (`auth`, `routes`, `db`) and all of them can be within a single module.

---

### **3. What is the Difference Between a Package and a Module?**

| Feature       | Package                 | Module                |
| ------------- | ----------------------- | --------------------- |
| Level         | Smaller Unit            | Larger Unit           |
| Contains      | Several Go Files        | Multiple Packages     |
| Config File   | None                    | Has `go.mod`          |
| Creation      | Automatic (by directory)| Via `go mod init`     |

🧠 Summary:

> Package = Small piece of code

> Module = Entire project containing multiple packages

---

### **4. Most Important Go CLI Commands**

* `go run` → Run a Go file
* `go build` → Build (compile) the program
* `go mod init` → Create a new module
* `go get` → Install a package from the internet
* `go fmt` → Format code files
* `go test` → Run tests
* `go doc` → Show documentation
* `go clean` → Clean up generated files

---

### **5. General Data Type Categories in Go**

Data in Go is divided into several groups:

* **Numeric** → int, float
* **String**
* **Boolean**
* **Composite** → array, slice, map, struct
* **Special** → pointer, interface, function

---

### **6. Golang Data Types (More Detailed)**

* Integers: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, ...
* Floating-Point: `float32`, `float64`
* Boolean: `bool`
* String: `string`
* Byte and Character: `byte`, `rune`
* Composite: `array`, `slice`, `map`, `struct`
* Others: `pointer`, `interface`, `function`

---

### **7. How is Enum Defined in Golang?**

Go doesn't have real enums, but they can be simulated using `const` and `iota`.

🧠 Example:

```go
type Day int

const (
  Sunday Day = iota
  Monday
  Tuesday
)
```
`iota` automatically assigns values 0, 1, 2, etc.

---

### **8. Definition of Pointer in Go and Its Usage**

A pointer is a **reference to a variable's memory address**.
Instead of holding the value, it holds its address.

🧠 Example:

```go
var x = 10
var p = &x // p points to the address of x
fmt.Println(*p) // Prints the value of x, which is 10
```

**Usage:**

* To change a variable's value within a function
* For better performance (avoid copying large data)

---

### **9. When Should We Use Pointers and When Should We Avoid Them?**

✅ **Use when:**

* A function needs to modify the original value
* Working with large data (e.g., a big struct)
* There's a need to share memory between different parts of the program

❌ **Avoid when:**

* Data is small and temporary
* Only reading the data is intended
* It causes unnecessary complexity in the code

---

### **10. What is the Definition of Rune in Go?**

`rune` is a data type used to **represent a Unicode character**.
In fact, `rune` is an alias for `int32`.

🧠 Example:

```go
var ch rune = 'A'
fmt.Println(ch)  // 65 (Unicode code point)
fmt.Printf("%c", ch) // A
```

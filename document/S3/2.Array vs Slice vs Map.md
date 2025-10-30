## 🧩 **۱. Definition of Composite Data Types in Go**

**Composite Data Types** are types that are composed of multiple simple (primitive) data elements.
They themselves contain several values of similar or different types.

In Go, these types include:

```go
array, slice, map, struct
```

🧠 Summary Definition:

> Types that hold multiple values or multiple fields within themselves, not just a single simple value.

For example:

- `array` → Multiple values of a specific type
- `slice` → Dynamic version of an array
- `map` → A collection of key-value pairs
- `struct` → A collection of fields with different types

---

## ⚖️ **۲. Differences Between Array, Slice, and Map in Go**

| Feature                    | **Array**                          | **Slice**                      | **Map**                                     |
| -------------------------- | ---------------------------------- | ------------------------------ | ------------------------------------------- |
| **Structure**              | Collection of fixed-size values    | Similar to array but dynamic size | Collection of key-value pairs               |
| **Size Determination**     | Specified and fixed at definition  | Automatically changeable        | Size automatically managed                  |
| **Data Types**             | All elements of the same type      | All elements of the same type  | Key and value can have different types      |
| **Dynamic (Mutable)**      | ❌ No                              | ✅ Yes                         | ✅ Yes                                      |
| **Indexing**               | By number (0,1,2,...)              | By number (0,1,2,...)          | By key                                      |
| **Common Usage**           | Fixed data with known size         | Dynamic and variable data       | Storing data with key-value relationships  |
| **Definition Example**     | `var arr [3]int = [3]int{1,2,3}`   | `nums := []int{1,2,3}`         | `m := map[string]int{"a":1,"b":2}`          |
| **Length (len)**           | `len(arr)` → Fixed                 | `len(slice)` → Variable        | `len(map)` → Number of keys                 |
| **Adding New Data**        | Impossible                         | With `append(slice, val)`      | With `m[key] = value`                       |

---

### 🔹 **Simple Summary for Memorization:**

> - **Array** → Fixed size

> - **Slice** → Dynamic version of Array

> - **Map** → Key and value

---

🧠 Short example of each:

```go
// Array
arr := [3]int{1, 2, 3}

// Slice
nums := []int{1, 2, 3}
nums = append(nums, 4)

// Map
m := map[string]int{
  "apple":  10,
  "banana": 5,
}
fmt.Println(m["apple"]) // 10
```

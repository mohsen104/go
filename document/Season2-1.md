## 🧩 انواع کلی دیتا تایپ‌ها در زبان Go

در Go داده‌ها به چند دسته‌ی اصلی تقسیم می‌شن:

---

### **1. Basic (یا Primitive Types) — نوع‌های پایه**

نوع‌هایی که مستقیماً مقدار رو نگه می‌دارن.
مثلاً اعداد، رشته، و بولین.

🧠 مثال‌ها:

```go
bool
string
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
complex64, complex128
byte (alias برای uint8)
rune (alias برای int32)
```

---

### **2. Composite Types — نوع‌های ترکیبی**

نوع‌هایی که از ترکیب چند داده‌ی ساده ساخته می‌شن.

🧠 مثال‌ها:

```go
array   // آرایه با اندازه ثابت
slice   // آرایه پویا
map     // کلید و مقدار (Key-Value)
struct  // ساختار داده‌ای شامل چند فیلد
```

---

### **3. Reference Types — نوع‌های ارجاعی**

به‌جای نگهداری مقدار، به آدرس حافظه اشاره می‌کنن.

🧠 مثال‌ها:

```go
pointer  // اشاره‌گر به یک مقدار
function // تابع به‌عنوان مقدار
channel  // برای ارتباط بین goroutine‌ها
interface // نوع انتزاعی برای چندین نوع مختلف
```

---

### **4. Interface Types — نوع‌های رابطی**

نوع‌هایی که فقط **رفتار (method)** تعریف می‌کنن، نه داده.
هر نوعی که اون متدها رو پیاده کنه، به اون interface تعلق داره.

🧠 مثال:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

---

### **5. Function Type — نوع تابع**

در Go توابع خودشون یک نوع داده هستن، یعنی می‌تونی تابع رو داخل متغیر بریزی یا به‌عنوان پارامتر بفرستی.

🧠 مثال:

```go
var f func(int) int
```

---

### 📘 خلاصه کلی جدول‌وار:

| دسته      | مثال                                  | توضیح             |
| --------- | ------------------------------------- | ----------------- |
| Basic     | int, bool, string                     | نوع‌های پایه      |
| Composite | array, slice, map, struct             | نوع‌های ترکیبی    |
| Reference | pointer, function, channel, interface | نوع‌های ارجاعی    |
| Interface | Writer, Reader                        | نوع‌های رفتاری    |
| Function  | func()                                | تابع به‌عنوان نوع |

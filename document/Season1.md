**1. What is the difference between Stack and Heap?**

*   **Stack** is used for temporary and small data (like local variables and function parameters).
    It is fast and managed automatically (it is cleared when the function ends).
*   **Heap** is used for data that has a longer lifespan (like dynamic data created with `new` or `make`).
    It is slower and requires memory management (in Go, this is handled by the Garbage Collector).

🧠 Summary:

> Stack = Fast, Automatic, Short-lived

> Heap = Slower, Flexible, Longer-lived

---

**2. What is the difference between JIT and AOT?**

*   **JIT (Just-In-Time)**: Compilation happens during runtime. Like JavaScript or Java, where code is converted to binary at runtime.
    🔹 Advantage: Flexible and can perform optimizations at runtime.
    🔹 Disadvantage: Initial execution is slightly slower.

*   **AOT (Ahead-Of-Time)**: Compilation happens before execution. Like Go or C, where the output is directly a binary file.
    🔹 Advantage: Faster execution and independent of an interpreter.
    🔹 Disadvantage: Less flexibility for runtime optimizations.

🧠 Summary:

> JIT = Compilation at runtime

> AOT = Compilation before execution
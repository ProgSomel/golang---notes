## 041 - Separate Stack For Separate Thread(A thread is the smallest unit of execution within a process.You can think of it as a lightweight process.)

**For each Thread there is one Stack**

- When we starts a program, then a process starts(Like: When we starts google chrome then a process starts)
  **Process** contains:
- Part of RAM-(code segment, data segment, stack, heap)
- Thread(if thread runs then process runs), if there is single thread than program and thread are same thing but if there is multiple thread then process and thread are different things, each thread works as a part of a process as lightweight process

> When process starts a initial thread starts <br>

## Let there is two threads in a process:

- First thread will be allocated to main stack
- For second thread, need another stack, so we have to tell that we need another stack(each stack needs 80mb(for linux) for each thread), stack can be anywhere on RAM

### Threads are not operated by processor, they are operated by OS, And OS are operated by CPU.

### Main part(core component) of operating system is Kernal

**------------------------------------------------------------------------------------------------------------------------------------**

# 🧵 "Separate Stack for Separate Thread"

## 📌 Basic Concepts First

### 🔹 Thread

- Thread হলো প্রোগ্রামের মধ্যে ছোট ছোট execution unit, যা OS-এ independent কাজ করে।
- একটি থ্রেড হলো process এর অংশ, যা তার নিজস্ব stack, registers, এবং program counter আছে।

### 🔹 Stack

- Stack হলো মেমোরির একটি বিশেষ অংশ যেখানে ফাংশন কল, return address, local variable ইত্যাদি রাখা হয়।
- Stack-এ data LIFO (Last-In, First-Out) হিসেবে রাখা হয়।

## 🌐 প্রত্যেক থ্রেডের জন্য আলাদা Stack কেন?

যখন একটি application-এ multiple threads চালানো হয়, তখন প্রতিটি থ্রেডের নিজস্ব Stack প্রয়োজন হয় কারণ:

- **Isolation:** প্রত্যেক থ্রেড নিজের ফাংশন কল ও local variable নিয়ে কাজ করতে পারে, অন্য থ্রেডকে interfere করবে না।
- **Concurrency:** একাধিক থ্রেড একসঙ্গে স্বাধীনভাবে কাজ করতে পারে।
- **Security and Safety:** এক থ্রেডের stack overflow হলে অন্য থ্রেডে প্রভাব পড়ে না।

## 🖥️ Memory Layout with Separate Stacks (Visual)

```less
Thread 1         Thread 2         Thread 3
--------         --------         --------
| funcC |        | funcX |        |        |
| funcB |        | funcY |        |        |
| funcA |        | main  |        | main   |
| main  |        |       |        |        |
--------         --------         --------
   ↑                ↑                ↑
  sp1              sp2              sp3
```


**sp (stack pointer):** Stack এর উপরের অংশকে নির্দেশ করে।

- প্রতিটি থ্রেডের stack pointer (sp) আলাদা থাকে।
- Stack মেমোরি প্রতিটি থ্রেডের জন্য OS allocate করে।

## 🚧 Stack-এ যা যা থাকে (Stack Frame)
প্রতিটি ফাংশন কলের সময় একটি **Stack Frame** তৈরি হয়, যাতে থাকে:

- **Return address** (ফাংশন থেকে কোথায় ফিরে যাবে)
- **Local variables** (int a, int b, ইত্যাদি)
- **Arguments passed** (ফাংশনে পাঠানো প্যারামিটার)
- **Registers/temporary** variables

উদাহরণ:
```c
func value(x int) {
    y := 10;
    // Local variable y, argument x
}
```
এর স্ট্যাক ফ্রেম দেখতে এরকম:

```sql
| y = 10     | <- local variable
| x          | <- argument
| Return addr| <- কোথায় ফিরবে
```

## 🌟 Stack vs Heap (কোথায় কী)

| বৈশিষ্ট্য        | Stack (Per thread)                          | Heap (Shared by threads)                     |
|------------------|---------------------------------------------|----------------------------------------------|
| **Size**         | সীমিত (প্রায় 1MB–8MB per thread)             | বড় (dynamic)                                 |
| **Lifetime**     | Function কল পর্যন্ত                          | যতক্ষণ manually free না করা হয়              |
| **Management**   | Automatic (OS handles)                       | Manual or via GC (Go, Java)                  |
| **Access**       | Thread-local                                 | All threads share                            |
| **Performance**  | Fast                                         | Slow (due to allocation overhead)            |


## 🔄 Thread Stack in Different Languages
### 🟠 C/C++ (OS Threads)

- Stack size fixed থাকে।
- খুব বেশি recursion করলে stack overflow হতে পারে।
- Stack size manual set করা যায় (pthread attributes দিয়ে):

```c
pthread_attr_t attr;
pthread_attr_init(&attr);
pthread_attr_setstacksize(&attr, 1<<20); // 1MB stack
```

### 🟢 Go (Goroutines)

- প্রতিটি Goroutine-এ growable stack থাকে।
- ছোট (~2KB) থেকে শুরু করে প্রয়োজনে নিজে থেকে বৃদ্ধি পায়।
- Stack overflow কম হয়, কারণ Go নিজেই stack size manage করে।

```go
func worker() {
    // প্রতিটি Goroutine নিজস্ব stack পাবে
    fmt.Println("Goroutine running")
}

func main() {
    go worker()
}
```

### 🔵 Java (JVM Threads)

- প্রতিটি থ্রেডে JVM একটি stack allocate করে।
- Default প্রায় 1MB থাকে, JVM flags দিয়ে modify করা যায়:

```diff
-Xss1m
```

## 🚩 Common Stack Issues
**⚠️ Stack Overflow**
- সাধারণত recursive function বা বড় local variable stack-এ রাখলে overflow হয়।

**⚠️ Dangling pointers (C/C++)**
- অন্য থ্রেডে স্ট্যাকের অ্যাড্রেস দিলে থ্রেড শেষ হলে অ্যাড্রেস invalid হয়ে যায়।

Unsafe Example (C):
```c
int *ptr;
void* worker(void* p){
    int local = 10;
    ptr = &local;
    return NULL; // local শেষ হলে ptr dangling হয়ে যায়
}
```

## 🚨 Thread Safety
- প্রত্যেক থ্রেডের stack আলাদা হওয়ায় stack local variable inherently thread-safe।
- কিন্তু heap এবং global variable share করা হলে mutex বা locks দিয়ে thread-safe করতে হয়।

## 🎯 Practical Benefits
- Multithreaded web servers (Apache, Nginx)
- Goroutine-based concurrency (Go)
- High-performance applications (games, graphics)

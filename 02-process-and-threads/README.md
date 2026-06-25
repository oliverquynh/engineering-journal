# 02. Process & Threads Fundamentals

## 📘 Part 1: Core Concepts

### 🏢 Program vs. Process
* **Program:** A static collection of files containing machine code, source code, and configurations stored on the disk. It remains dormant until executed.
* **Process:** A program in execution. In other words, it is an active instance of a program currently being processed by the CPU and allocated resources by the Operating System.

### 🚦 Common Process States
* **Ready:** The process is loaded and waiting for CPU time to be allocated.
* **Running:** The process's instructions are currently being executed by the CPU.
* **Waiting / Blocked:** The process is paused, waiting for an external event or resource (e.g., waiting for user keyboard input).
* **Zombie:** A process that has completed execution, but its entry still remains in the process table because its parent process hasn't read its exit status yet.
* **Orphan:** A child process whose parent process has terminated or died, leaving it to be adopted by the `init` (or `systemd`) process.
* **Terminated:** The process has finished its execution and released its resources.

---

### 🧵 Single-Threading vs. Multi-Threading
A **Thread** is the smallest unit of execution within a process. A process can contain one or multiple threads.

* **Single-Threaded:** A process that has only one thread of execution (the `main` thread). Tasks are executed sequentially (Task B can only start after Task A finishes).
* **Multi-Threaded:** A process that utilizes multiple threads. Instructions can be executed via **Concurrency** or **Parallelism**.

#### 🔄 Concurrency vs. Parallelism
* **Concurrency:** Threads take turns running on a **single CPU core** at an incredibly high speed. The OS constantly switches back and forth between threads via **Context Switching**, creating the illusion that they are running simultaneously.
* **Parallelism:** Threads literally execute at the **exact same physical moment** on **multiple CPU cores** (e.g., an 8-core CPU).

When multiple parallel threads access shared resources (RAM, Disk, etc.) simultaneously, two classic concurrency issues can occur:
* **Race Condition:** Occurs when two or more threads write to the same memory location at the same time, leading to data corruption.
    * *Real-world Analogy:* An e-commerce store has only 1 item left in stock. Two customers place an order at the exact same millisecond. Both orders succeed, forcing the shop owner to cancel one order and apologize to the customer.
* **Deadlock:** Occurs when two or more threads are blocked forever, each waiting for a resource held by the other, causing the system to freeze.
    * *Real-world Analogy:* Worker A is holding a Hammer and needs a Pair of Pliers to proceed. Worker B is holding the Pliers and needs the Hammer. Both refuse to give up their tools and stand there staring at each other. Progress grinds to a halt.

---

### 🏭 The Factory Analogy
To easily visualize how processes and threads operate: **Think of a Process as a Factory, and a Thread as a Worker.**

* The **Factory (Process)** provides the physical workspace, tools, and raw materials (CPU, RAM, and system resources allocated by the OS).
* The **Workers (Threads)** share the exact same workspace and tools within that factory.
* **Single-threading** is like a factory with only **one worker** who has to do everything from A to Z sequentially.
* **Multi-threading** is like a factory with **multiple workers**, each handling a specific sub-task to maximize productivity.
    * **Concurrency (The Multi-tasking Worker):** A single worker operates multiple machines. Instead of waiting for Machine 1 to finish heating up, they move over to turn on Machine 2, then jump back to check Machine 1. They switch tasks so fast that it looks like they are running all machines at once.
    * **Parallelism (The Master Switch):** The factory has a master control panel. When pressed, all machines are powered up and run independently at the exact same physical moment.

---

## 💻 Part 2: Hands-On Laboratory

### 🔍 Example 1: Inspecting a Live Process
Within a process, how can we answer these questions?
- What is its ID?
- What is its state?
- How much CPU and RAM does it hold?
- How many threads does the OS create for it?

#### Steps to Execute:
1. Start a process (an HTTP server):
    ```bash
    go run server.go
    ```
2. Look for the created process:
    ```bash
    ps aux | grep "go run server.go"
    ```
3. You will see a line similar to this:
    ```bash
    quynhnx    16360  0.0  0.1 1602236 12580 pts/5   Sl+  11:24   0:00 go run server.go
    ```

#### Key Metrics to Notice:
* **Process ID (PID):** `16360`
* **CPU Usage:** `0.0%` (very tiny)
* **Memory Usage:** `0.1%`
* **Process State (`Sl+`):**
    * `S`: Interruptible Sleep
    * `l`: Multi-threaded
    * `+`: Running in the foreground
* To check how many threads the OS created for this specific process, run:
    ```bash
    ls /proc/16360/task | wc -l
    ```

---

### ⚡ Example 2: Simulating a Race Condition
In this program, 5 threads concurrently read and write to the same `counter` variable (the same memory address in RAM).

For example, thread 1 and thread 2 might simultaneously read the counter value as `10`. Both will increment it and write `11` back to RAM, instead of the expected `12`. This data loss is the classic symptom of a Race Condition.

Because thread scheduling is managed by the OS and executed by the CPU, the final value of the counter changes unpredictably with each run.

#### Run the Code:
```bash
go run race.go
```

#### Execution Results:

```plain
Starting 5 concurrent threads, each incrementing 100000 times...
Expected Theoretical Result: 500000
--------------------------------------------------
ACTUAL RESULT IN RAM: 304800
⚠️ DATA CORRUPTION DETECTED! Lost: 195200 units.

Starting 5 concurrent threads, each incrementing 100000 times...
Expected Theoretical Result: 500000
--------------------------------------------------
ACTUAL RESULT IN RAM: 341960
⚠️ DATA CORRUPTION DETECTED! Lost: 158040 units.

Starting 5 concurrent threads, each incrementing 100000 times...
Expected Theoretical Result: 500000
--------------------------------------------------
ACTUAL RESULT IN RAM: 269792
⚠️ DATA CORRUPTION DETECTED! Lost: 230208 units.
```

## 🌐 Part 3: Real-World Use Cases

### 🔒 Distributed Locks in Laravel (Redis)

In large-scale distributed systems, Race Conditions do not just happen between threads inside a single RAM space; they also occur between independent background workers competing for global resources (e.g., preventing duplicate deployment jobs on a remote server).

Laravel solves this by using Atomic Locks backed by Redis:

1. The Scenario: Two separate workers trigger a deployment job at almost the exact same millisecond.

2. The Atomic Operation: The first job to hit Redis executes an atomic command (SETNX). Redis checks for the lock and creates it in a single, un-interruptible step.

3. The Race Resolved: Even if the second job arrives just 0.1ms later, Redis detects the existing lock and rejects the request. The second job safely aborts or reschedules itself, preventing server configuration corruption.

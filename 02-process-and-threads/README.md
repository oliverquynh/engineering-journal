## Example 1

Within a process, how can we answer these questions?

- What is its ID?
- What is its state?
- How much CPU and RAM does it hold?
- How much threads does OS create for it?

### Hands-on

- Start a process. It's an HTTP server.
    ```bash
    go run server.go
    ```

- Look for the created process.

    ```bash
    ps aux | grep "go run server.go"
    ```

- You will see a line looks similar to this.

    ```bash
    quynhnx    16360  0.0  0.1 1602236 12580 pts/5   Sl+  11:24   0:00 go run server.go
    ```
- What you need to pay attension are:
    - The process ID is 16360
    - CPU usage is 0.0% (it's really tiny)
    - MEM usage is 0.1%
    - Sl+ means this process is:
        * S: Interruptible Sleep
        * l: Multi-threaded
        * +: Running in the foreground
- To know how many threads OS creates for it, try running
    ```bash
    ls /proc/16360/task | wc -l
    ```

## Example 2

In this program, 5 threads concurrently read and write to the same counter variable (the same memory address in RAM).

For example, thread 1 and thread 2 might simultaneously read the counter value as 10. Both will increment it and write 11 back to RAM, instead of the expected 12. This data loss is the classic symptom of a Race Condition.

Because thread scheduling is managed by the OS and executed by the CPU, the final value of counter changes unpredictably with each run.

Running the Code

```bash
go run race.go
```

Execution Results:

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

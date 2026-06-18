## Example 1

Within a process, how can we answer these questions?

- What is its ID?
- What is its state?
- How much CPU and RAM does it hold?
- How much threads does OS create for it?

### Hands-on

- Start a process. It's an HTTP server.
    ```bash
    ./server
    ```

- Look for the created process.

    ```bash
    ps aux | grep "\.\/server"
    ```

- You will see a line looks similar to this.

    ```bash
    quynhnx    16360  0.0  0.1 1602236 12580 pts/5   Sl+  11:24   0:00 ./server
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

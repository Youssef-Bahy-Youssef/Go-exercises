> [!NOTE] Exercise 7.2
> 
> Write a function `CountingWriter` with the following signature:
>
> ```go
> func CountingWriter(w io.Writer) (io.Writer, *int64)
> ```
>
> Given an `io.Writer`, it should return a new `Writer` that **wraps** the original and a pointer to an `int64` variable that always contains the number of bytes written to the new Writer.

## Solution

To implement `CountingWriter`:

1. **Create a new struct** to wrap the original writer and maintain the byte count:

   - The first field is the `io.Writer` to delegate actual writes.
   - The second field is an `int64` variable to keep track of bytes written.

2. **Implement the `Write` method** for the new type:

   - Forward the byte slice to the wrapped writer.
   - Update the internal `count` variable with the number of bytes written.
   - Return the number of bytes written and any error from the underlying writer.

3. **Implement the `CountingWriter` function**:

   - Instantiate the new struct, wrapping the provided `io.Writer`.
   - Return the struct as an `io.Writer` and a pointer to its `count` variable.

This approach ensures that:

- All writes are counted automatically.
- The returned pointer always reflects the current number of bytes written.
- The original writer’s behavior is preserved via delegation.
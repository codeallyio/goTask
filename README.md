# Go Interview Task

## Background



## TODO

All you can change is `main.go` file. You are not allowed to change anything in `mockserver.go`.

1. Implement `isActive()` `connection`'s method. It should validate if the lifespan of the connection did not expire.

2. You have implement the rest of the `main` function. It's current code has to stay in place. Being more specific, you have to handle all requests from the activeRequest channel. `handleRequest` function should do the job.

```
func (c *connection) isActive() bool {
  << YOUR CODE GOES HERE >>
}

func main() {
  runMockServer()
  defer checkSuccess()
  
  << YOUR CODE GOES HERE >>
}
```

Make your best to handle all 9 requests:
```
=== Success ===
Completed 9 out of 9 requests
```
You can write additional functions if you wish, if only they sit in `main.go` file

### Good luck!

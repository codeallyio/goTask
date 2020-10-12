# Go Interview Task

## Background

There is a mock server inside `mockserver.go` file. To run the mock server, there is a `runMockServer()` function. It will create connections using pre-defined `mockConnections` slice of `connection`s. After the connections are being assigned to `activeConnections`, this function will run `generateSomeRequests`. As the name suggests, it will generate requests, you will have to handle. They will sit into `activeRequests` channel. 

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

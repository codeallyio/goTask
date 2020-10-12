// GO CONCURRENCY TASK

package main

// Connection's holdUntil time is being set to 3 seconds when getting connected.
// Implement function, checking if connection's holdUntil time exceeded current time.
func (c *connection) isActive() bool {
	return false
}

func main() {
	runMockServer()
	defer checkSuccess()

	// Here handle all the requests from activeRequest channel.
	// For this use the `handleRequest(r request)` function.
}

# Go HTTP Request Body Resource Leak

This repository demonstrates a common error in Go when handling HTTP requests: forgetting to close the request body.  This can lead to resource leaks, especially when dealing with large requests.

## The Bug (`bug.go`)

The `bug.go` file shows how a request body can be left open, leading to a potential resource leak.  It omits the crucial `defer r.Body.Close()` call.

## The Solution (`bugSolution.go`)

The `bugSolution.go` file demonstrates the correct way to handle the request body.  By using `defer r.Body.Close()`, the body is guaranteed to be closed after the function completes, even if errors occur.

## How to Run

1. Clone this repository
2. Navigate to the directory
3. Run the `bug.go` file (observe the error).
4. Run the `bugSolution.go` file (observe the proper behavior).

This example highlights the importance of diligent resource management in Go to avoid potential problems in production systems.
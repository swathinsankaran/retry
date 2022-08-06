# retry

A simple Go retry package and it is heavily inspired from the [try](https://github.com/matryer/try) package from [Mat Ryer](
https://github.com/matryer/try)

```
go get github.com/swathinsankaran/retry
```

### Usage

`retry.Do` with the function you want to retry along with number of retries(`maxRetryAttempt`) and a timeout for the entire execution(`retryTimeout`).

The `retry` package retries when it encounters any of the below events:

  * Function passed returns an error. 
  * Execution of the function passed was unsuccessful. 

The `retry` package returns error in the below cases:
  * Function passed returns an error.
  * Retry count exceeds `maxRetryAttempt`.
  * Function execution time exceeds `retryTimeout` milliseconds.

Example:
```
import "github.com/swathinsankaran/retry"

// some codes...

err := retry.Do(func() (bool, error) {
    // code to retry goes here...
}, 10, 1)
```

In the above example the function will be called repeatedly until:
  * error is `nil`.
  * retries exceeds `10`.
  * execution time exceeds `1 millisecond`.
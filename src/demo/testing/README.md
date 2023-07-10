# Testing in Go

Import to test software to prevent regressions and ensure it meets specifications

- Unit testing - test individual functions
- Integration testing - test functions/modules working together

Go makes no distinction between the two
- Same process to create both Unit & Integration test

## Setup test

Tests are written in separate files, sharing the name of the file they are testing eg:
- `fileNamePkg.go` -> `fileNamePkg_test.go`
<br>
- Unit tests should be in the same package
- The testing package is used to create tests and must be imported in each test file

## Example testing function
Function File: `validateEmail.go`

```go
package main
import "regexp"

func IsValidEmail(address string) bool {
  re, ok := regexp.Compile(`.+@.+\..+`)
  if ok != nil {
    panic("Failed to compile regex")
  } else {
    return re.Match([]byte(address))
  }
}
```

Test File: `validateEmail_test.go`

```go
package main
import "testing"  // import testing package

func TestIsValidEmail(test *.testing.T) {
  data := "test@example.com"
  if !IsValidEmail(data) {
    test.Errorf("IsValidEmail(%v)=false, expected to true", data)
  }
}
```

## Running test

Execute `go test` to run your tests
Run `go test` command in your terminal

## Available Testing Functions

Many testing functions available in the **testing** package
- `Fail()` - Mark the test as failed
  - `Errorf(string)` - Fail & add a message
- `FailNow()` - Mark the test as failed, abort current test
  - `Fatalf(string)` - Fail, abort, and add a message
- `Logf()` - Equivalent to `Printf`, but only when test fails

## Test Tables
- Usually need to test more than one set of input data
- **Test tables** can be used to accomplish this
  - Similar to parameterized testing

### Testing Tables: Example

```go
func TestIsValidEmailTable(test *testing.T) {
  table := []struct {
    email string
    expect bool
  }{
    {"test@example.com", true},
    {"test@fail", false},
    // ...
  }

  for _, data := range table {
    result := IsValidEmail(data.email)
    if result != data.expect {
      test.Errorf("%v: %t, want: %t", data.email, result, data.expect)
    }
  }
}
```

### Testing `queue.go` file

1. Go to `cd testing/`
2. Run `go test -v .`

On your terminal should appear testing result:

```sh
go test -v .

=== RUN   TestAddQueue
--- PASS: TestAddQueue (0.00s)
=== RUN   TestNext
--- PASS: TestNext (0.00s)
PASS
ok  	go-fundamental/src/demo/testing	(cached)
```

## Summary

- Test files end with `_test`
- Use the testing package to perform tests
  - Tests are ran with the command `go test`
- Test tables can be used to test multiple pieces of data
- Use `t.Errorf` to report an error
- Use `t.Fatalf` to report an error, and also abort the test case
- Use `t.Logf` to display debug or test messages

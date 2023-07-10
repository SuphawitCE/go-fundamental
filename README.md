# Go Fundamental

Go Fundamental is a project to learn Go programming, which aims to gain knowledge of language features, basic syntax, best practices, pros, and cons, and to be practical in production. study on the Go lesson, write some code demos, and apply the lesson to solve the problem.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have a Mac / Linux Workstation.
- You have installed the following software:
  - Go version `1.20`

## Using `Go Fundamental` project on local

```sh
git clone git@github.com:SuphawitCE/go-fundamental.git
cd go-fundamental/
```

## Project structure

- `src/`
  - `src/demo`: Go lesson and playground to demonstrate the basic and theory
  - `src/problems`: Problem based on lesson

### Conventional commit
`<type>`(`<scope>`): `<description>`

Since our JIRA does not yet integrate with our GitHub, it is needless to prefix your commit with JIRA's ticket.


#### Type

- _fix_: a commit of the type fix patches a bug in your codebase (this correlates with PATCH in semantic versioning).

- _feat_: a commit of the type feat introduces a new feature to the codebase (this correlates with MINOR in semantic
  versioning).

- _refactor_: a commit of the type refactor organizes a codebase without changing its logic.

- _chore_: a commit of the type chore is the change that does not affect the meaning of the code (white-space,
  formatting, missing semi-colons, comments, etc)

- _docs_: a commit of the type docs is for a documentation only changes i.e., update readme file.

- _ci_: a commit of type ci changes CI configuration files and scripts.

- _build_: a commit of the type build changes the build system or external dependencies (example scopes: gulp, broccoli,
  npm).

- _test_: a commit of the type test adds the missing tests or correct the existing tests.


# Further reading
- [Go](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Tour of Go](https://go.dev/tour/list)
- [Testing package](https://pkg.go.dev/testing)
- [Go By Example](https://gobyexample.com/)
- [Comprehensive Guide to Testing in Go](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)
- [Go 101](https://go101.org/)
- [Dave Cheney](https://dave.cheney.net/practical-go)

# Contributing

Read for following docs to get started:

- Go tutorial - https://gobyexample.com/hello-world
- cobra (CLI stuff) - https://github.com/spf13/cobra
- bubbletea (TUI stuff) - https://github.com/charmbracelet/bubbletea

## Setting up

After cloning this repository...

Run:

```
go run .
```

Build (produces `pompup` binary):

```
go build
```

## Etc

- use conventional commits
- write idiomatic Go
- all `exec.Command` must be in the following form:
  ```go
  cmd := exec.Command("some", "commands", "and", "args")
  cmd.Stderr = os.Stderr // important
  // ...
  cmd.Run()
  ```

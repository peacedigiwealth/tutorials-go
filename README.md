# tutorials-go

### Created a new repository in my Github account for Tutorials-go

    https://github.com/peacedigiwealth/tutorials-go.git

#   Create the go.mod file having:

    module github.com/peacedigiwealth/tutorials-go

    go 1.20

## Adding packages and modules

Short guidance for contributors on how to add new code to this repo:

- Prefer a single top-level module: add packages as subdirectories (for example `examples/`, `pkg/`, or `cmd/`).
    - Example package path: `github.com/peacedigiwealth/tutorials-go/examples/hello`

- Add new Go files under a subdirectory and follow standard package/import rules.
    - Example of a tiny program:
        ```go
        // examples/hello/main.go
        package main

        import "fmt"

        func main() { fmt.Println("hello") }
        ```

- After adding or changing imports, run:
    ```bash
    go mod tidy   # updates go.sum and downloads any new deps
    go test ./... # run tests
    ```

- When to create a sub-module (`go.mod` in a subfolder):
    - Create a sub-module only when the subproject needs independent versioning or release cadence.
    - Initialize a sub-module with `go mod init <module-path>` inside that subfolder and consider adding a `go.work` at repository root for local development across modules.

- Versioning note: if you publish a v2+ major version for any module, include the major in the module path (for example `module github.com/peacedigiwealth/tutorials-go/v2`).

- Commit `go.mod` and `go.sum` files in the same PR as code changes so CI and other developers get consistent dependency information.

If you'd like, I can add a small example that imports a third-party package to demonstrate `go mod tidy` creating `go.sum`.


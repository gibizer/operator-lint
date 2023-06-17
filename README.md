[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/gibizer/operator-lint/main.svg)](https://results.pre-commit.ci/latest/github/gibizer/operator-lint/main)
# operator-lint

Static analysis library for k8s operators created by [operator-sdk](https://sdk.operatorframework.io/)

## Usage
```bash
  go install github.com/gibizer/operator-lint/
  go vet -vettool=$(which operator-lint)
```

## Checks
| Check | Category | Added in | Description
|---|---|---|---|
| [T001](linters/envtest/T001) | EnvTest | v0.1.0 | checks that Gomega's `Eventually` and `Consistently` blocks use a local Gomega instance for asserts
|---|---|---|---|
| [C001](linters/crd/C001) | CRD | v0.1.0 | detects incompatible `Required` and `Optional` kubebuilder markers
| [C002](linters/crd/C002) | CRD | v0.1.0 | detects incompatible `Required` kubebuilder marker and `omitemty` golang tag
| [C003](linters/crd/C003) | CRD | v0.2.2 | detects incompatible defaulting via `Optional` kubebuilder marker and `omitemty` golang tag

### Flags
Command line flags can be used to configure each check in the form of:
```bash
go vet -vettool=$(which operator-lint) -<check name>.<flag>
```

Common flags supported by each check:
* `<check name>.skip`: can be used to disable the given check. E.g. the
  following command will not execute the check C003:
  ```bash
  go vet -vettool=$(which operator-lint) -C003.skip
  ```

You can use `go vet -vettool=$(which operator-lint) -flags` to print out all
the supported flags

## Adding a new check
- Use `make new-lint` to generate a new empty linter under `linters`.
- Update the [README.md](README.md) with the description of the check
- If the test data for the check has dependencies then
  - Those dependencies needs to be imported in
    [force_test_deps.go](force_test_deps.go)
  - The [vendor](vendor) directory needs to be symlinked to the directory
    holding the test data
  - Need to run
    ```shell
    go get <dep>
    go mod tidy
    go mod vendor
    ```
This whole dance is needed as the
[analysistest package](https://pkg.go.dev/golang.org/x/tools/go/analysis/analysistest)
does not support Go Modules currently.

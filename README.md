[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/gibizer/operator-lint/main.svg)](https://results.pre-commit.ci/latest/github/gibizer/operator-lint/main)
# operator-lint

Static analysis library for k8s operators created by [operator-sdk](https://sdk.operatorframework.io/)

## Usage
```bash
  go install github.com/gibizer/operator-lint/
  go vet -vettool=$(which operator-lint)
```

## Checks
| Check | Category | Description
|---|---|---|
| [T001](linters/envtest/T001) | EnvTest | checks that Gomega's `Eventually` and `Consistently` blocks use a local Gomega instance for asserts
|---|---|---|
| [C001](linters/crd/C001) | CRD | checks incompatible kubebuilder markers and golang tags on CRD fields


## Adding a new check
- Add a new package under [linters](linters) for the new check
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

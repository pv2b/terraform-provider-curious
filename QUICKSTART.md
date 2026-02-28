# Quick Start Guide

## Getting Started

### 1. Initialize Go Module

First, update the module path in `go.mod` to match your GitHub username or organization:

```bash
# Edit go.mod and replace "yourusername" with your actual username
sed -i 's/yourusername/YOUR_GITHUB_USERNAME/g' go.mod main.go internal/provider/provider.go
```

### 2. Download Dependencies

```bash
go mod tidy
```

### 3. Build the Provider

```bash
make build
```

Or manually:

```bash
go build -o terraform-provider-curious
```

### 4. Install Locally for Testing

```bash
make install
```

This will install the provider to your local Terraform plugins directory.

### 5. Test the Provider

Create a test Terraform configuration:

```bash
cd examples
terraform init
terraform plan
```

You should see the outputs from the example functions.

### 6. Run Tests

Run unit tests:

```bash
make test
```

Run acceptance tests (requires Terraform >= 1.8):

```bash
make testacc
```

## Available Functions

| Function | Description | Example |
|----------|-------------|---------|
| `ascii` | Remove non-ASCII characters | `provider::curious::ascii("Hello 世界!")` → `"Hello !"` |
| `ascii_printable` | Keep printable ASCII only | `provider::curious::ascii_printable("Hi\t\n!")` → `"Hi!"` |
| `latinize` | Remove diacritics | `provider::curious::latinize("räksmörgås")` → `"raksmorgas"` |
| `kebab` | Convert to kebab-case | `provider::curious::kebab("Hello World")` → `"hello-world"` |
| `camel` | Convert to camelCase | `provider::curious::camel("Hello World")` → `"helloWorld"` |
| `snake` | Convert to snake_case | `provider::curious::snake("Hello World")` → `"hello_world"` |
| `elite` | Consonants upper, vowels lower | `provider::curious::elite("Sponge Bob")` → `"SPoNGe BoB"` |
| `sponge` | Alternate lower/upper | `provider::curious::sponge("sponge bob")` → `"sPoNgE bOb"` |

## Next Steps

1. **Customize the functions**: Add more normalization functions in [internal/provider/functions.go](internal/provider/functions.go)
2. **Update tests**: Add tests for new functions in [internal/provider/functions_test.go](internal/provider/functions_test.go)
3. **Configure release**: Update `.goreleaser.yml` with your GPG key for signing releases
4. **Publish**: Push to GitHub and create releases using GoReleaser

## Publishing to Terraform Registry

To publish your provider to the Terraform Registry:

1. Sign up at https://registry.terraform.io/
2. Link your GitHub repository
3. Create a GPG key for signing releases
4. Tag a release (e.g., `v0.1.0`)
5. Use GoReleaser to create the release:

```bash
export GPG_FINGERPRINT=your_gpg_fingerprint
git tag v0.1.0
git push origin v0.1.0
goreleaser release --clean
```

## Development Workflow

1. Make changes to the provider code
2. Run `make fmt` to format code
3. Run `make test` to run tests
4. Run `make install` to install locally
5. Test with Terraform configuration
6. Commit and push changes

## Debugging

To run the provider with debugging support:

```bash
go build -o terraform-provider-curious
./terraform-provider-curious -debug
```

Then use the `TF_REATTACH_PROVIDERS` environment variable as instructed in the output.

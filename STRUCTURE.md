# Project Structure

```
tf-curious/
├── .github/
│   └── workflows/
│       ├── test.yml              # CI workflow for testing
│       └── release.yml           # CD workflow for releases
├── examples/
│   └── functions.tf              # Example Terraform configurations
├── internal/
│   └── provider/
│       ├── provider.go           # Main provider implementation
│       ├── provider_test.go      # Provider test setup
│       ├── functions.go          # All function implementations
│       └── functions_test.go     # Function tests
├── .gitignore                    # Git ignore rules
├── .goreleaser.yml               # GoReleaser configuration
├── go.mod                        # Go module definition
├── main.go                       # Provider entry point
├── Makefile                      # Build automation
├── README.md                     # Main documentation
├── QUICKSTART.md                 # Quick start guide
└── terraform-registry-manifest.json  # Terraform Registry metadata

## Key Files

### Provider Core
- **main.go**: Provider entry point that starts the plugin server
- **internal/provider/provider.go**: Provider definition and configuration
- **internal/provider/functions.go**: Implementation of all normalization functions

### Functions Implemented
1. `ascii` - Removes non-ASCII characters
2. `ascii_printable` - Keeps printable ASCII only
3. `latinize` - Removes diacritics
4. `flat` - Converts to flatcase
5. `kebab` - Converts to kebab-case
6. `camel` - Converts to camelCase
7. `pascal` - Converts to PascalCase
8. `snake` - Converts to snake_case
9. `upper` - Converts to UPPER_CASE
10. `train` - Converts to TRAIN-CASE
11. `ada` - Converts to Ada_Case
12. `elite` - Consonants upper, vowels lower
13. `sponge` - Alternating lower/upper

### Build & Development
- **Makefile**: Provides convenient commands (build, install, test, etc.)
- **.goreleaser.yml**: Automated release configuration
- **.github/workflows/**: CI/CD automation

### Documentation
- **README.md**: Complete usage documentation
- **QUICKSTART.md**: Step-by-step getting started guide
- **examples/functions.tf**: Working examples of all functions

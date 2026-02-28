# Terraform Provider: Curious

```
   ______      ____  _             _____
  / ____/_  __/ __ \(_)___  __  __/ ___/
 / /   / / / / /_/ / / __ \/ / / /\__ \ 
/ /___/ /_/ / _, _/ / /_/ / /_/ /___/ / 
\____/\__,_/_/ |_/_/\____/\__,_//____/  
```

The curious case of a Terraform provider that exposes functions for string normalization, latinication and case transformation. This provider does not manage any resources or data sources - it only provides utility functions for string manipulation.

## Features

This provider offers the following custom functions:

- **`ascii`**: Removes diacritics first (latinizes), then removes all non-ASCII characters from a string (keeps 0-127)
- **`ascii_printable`**: Removes diacritics first (latinizes), then keeps only printable ASCII characters (32-126), excluding control characters like tabs and newlines
- **`latinize`**: Removes diacritics (accents) from strings, converting accented characters to their base Latin equivalents

**Case Conversion Functions:**
- **`flat`**: Converts to flatcase (all lowercase, no separators)
- **`kebab`**: Converts to kebab-case (lowercase with hyphens)
- **`camel`**: Converts to camelCase (first word lowercase, rest capitalized)  
- **`pascal`**: Converts to PascalCase (all words capitalized)
- **`snake`**: Converts to snake_case (lowercase with underscores)
- **`upper`**: Converts to UPPER_CASE (uppercase with underscores)
- **`train`**: Converts to TRAIN-CASE (uppercase with hyphens)
- **`ada`**: Converts to Ada_Case (capitalized words with underscores)
- **`elite`**: uPPeRCaSeS CoNSoNaNTS aND LoWeRCaSeS VoWeLS, TReaTiNG LeTTeRS WiTH DiaCRiTiCS aS VoWeLS
- **`sponge`**: aLtErNaTeS lOwEr/uPpEr cAsE oN lEtTeRs, sTaRtInG wItH lOwErCaSe

All case conversion functions latinize input first except `elite` and `sponge`. The word-based formats split on non-alphanumeric characters, while `elite` and `sponge` preserve non-letters.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.8
- [Go](https://golang.org/doc/install) >= 1.22

OpenTofu should work too, of course.

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using:

```sh
go build -o terraform-provider-curious
```

## Installing The Provider

### For Development/Testing

To install the provider locally for testing:

1. Build the provider (see above)
2. Create a local provider configuration:

```sh
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/pv2b/curious/0.1.0/linux_amd64
cp terraform-provider-curious ~/.terraform.d/plugins/registry.terraform.io/pv2b/curious/0.1.0/linux_amd64/
```

Adjust the path based on your OS and architecture. Or don't. It probably will break if you don't, but I
honestly have no idea.

## Using the Provider

### Configuration

The provider does not require any configuration:

```terraform
terraform {
  required_providers {
    curious = {
      source = "pv2b/curious"
      version = "~> 0.1"
    }
  }
}

provider "curious" {
  # No configuration required
}
```

### Function Examples

#### ASCII Only

```terraform
locals {
  original = "Hello, 世界! This is a test 🚀 with Café"
  cleaned  = provider::curious::ascii(local.original)
}

output "ascii_only" {
  value = local.cleaned
  # Output: "Hello, ! This is a test  with Cafe"
}
```

#### Normalize to Alphanumeric + Hyphens (using Terraform's replace)

```terraform
locals {
  user_input = "My-Service 123!@# (Prod) Café"
  # Latinize first to remove diacritics, then use regex to keep only alphanumeric and hyphens
  normalized = replace(provider::curious::latinize(local.user_input), "/[^A-Za-z0-9-]/", "")
}

output "normalized_id" {
  value = local.normalized
  # Output: "My-Service123ProdCafe"
}
```

#### Remove Diacritics (Latinize)

```terraform
locals {
  swedish_text = "räksmörgås"
  latinized    = provider::curious::latinize(local.swedish_text)
}

output "latinized" {
  value = local.latinized
  # Output: "raksmorgas"
}
```

#### Chaining Functions

You can chain multiple functions together with Terraform's built-in functions:

```terraform
locals {
  raw_input = "Café 世界! How are YOU? 🎉"
  
  # Latinize, then use regex to keep only alphanumeric and hyphens, then lowercase
  clean_id = lower(
    replace(
      provider::curious::latinize(local.raw_input),
      "/[^A-Za-z0-9-]/",
      ""
    )
  )
}

output "clean_id" {
  value = local.clean_id
  # Output: "cafhowareyou"
}
```

## Development

### Testing

Run the acceptance tests:

```sh
TF_ACC=1 go test ./... -v -timeout 120m
```

For the record I have no idea what acceptance tests are, and why the LLM that generated the slop above seems to think a 2 hour timeout is reasonable for testing some string manipulation, but the code seems to work, so I hereby accept it.

### Generating Documentation

If using tfplugindocs:

```sh
go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest
tfplugindocs generate
```

That didn't work when I tried it, something about tfplugindocs not appearing on my PATH when I tried to use it. Runnning it like this seems to work around that issue:

```sh
$(go env GOPATH)/bin/tfplugindocs generate
```

... or you could probably set up your PATH properly, idk.

## GReeTZ

- Romain Barissat (politician on GitHub) for his work on https://github.com/Olivr/terraform-null-normalize. It's a really clever hack to do latinization and normalization in pure Terraform, but writing a whole module declaration just to normalize a string, but now that Terraform providers can expose functions, this is a cleaner way of doing things while not polluting state files with string transformations. None of his code was used in this project (I'm using libraries from the golang x-repositories for that. I don't know what x-repositories are. The truth is out there, but it's not here.)
- The authors and maintainers of the libraries used in this project, including:
  - HashiCorp Terraform Plugin Framework and related Terraform SDK components
  - HashiCorp Terraform Plugin Testing
  - golang.org/x/text (Unicode normalization)
- Github Copilot. I have no idea how to code in Golang (specifically) or how to write a Terraform provider. Maybe I'll learn in the future. Until then, let it be known this project was shamelessly vibecoded, and this "project" (such as it is) as such stands - nay - tramples on the shoulders of giants past who have taught it everything it pretends to know.

## License

This provider is released into the public domain. See LICENSE for more information. I'm not the author of any of this code by any stretch of the imagination. All I did was write some prompts and add some terrible cringe into the README after the AI minions finished their work.
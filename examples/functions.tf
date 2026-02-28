terraform {
  required_providers {
    curious = {
      source = "registry.terraform.io/pv2b/curious"
    }
  }
}

provider "curious" {
  # No configuration required
}

# Example 1: Remove non-ASCII characters
locals {
  text_with_unicode = "Hello, 世界! This is a test 🚀 with émojis"
  ascii_only       = provider::curious::ascii(local.text_with_unicode)
}

output "ascii_only" {
  value       = local.ascii_only
  description = "String with only ASCII characters"
}

# Example 2: Normalize to alphanumeric and hyphens using Terraform built-ins
locals {
  messy_input  = "My-Service 123!@# (Prod) 🚀"
  # Latinize first, then use regex to keep only alphanumeric and hyphens
  clean_id     = replace(provider::curious::latinize(local.messy_input), "/[^A-Za-z0-9-]/", "")
}

output "normalized_id" {
  value       = local.clean_id
  description = "String with only alphanumeric characters and hyphens"
}

# Example 3: Remove diacritics (latinize)
locals {
  swedish_text  = "räksmörgås"
  latvian_text  = "Laima, Dievs, Sūņa"
  french_text   = "Café résumé naïve"
  
  latinized_swedish = provider::curious::latinize(local.swedish_text)
  latinized_latvian = provider::curious::latinize(local.latvian_text)
  latinized_french  = provider::curious::latinize(local.french_text)
}

output "latinized_swedish" {
  value       = local.latinized_swedish
  description = "Swedish text with diacritics removed: 'raksmorgas'"
}

output "latinized_latvian" {
  value       = local.latinized_latvian
  description = "Latvian text with diacritics removed"
}

output "latinized_french" {
  value       = local.latinized_french
  description = "French text with diacritics removed: 'Cafe resume naive'"
}

# Example 4: Remove non-printable ASCII characters
locals {
  text_with_controls = "Hello\tWorld!\nLine 2 End"
  printable_only     = provider::curious::ascii_printable(local.text_with_controls)
}

output "printable_only" {
  value       = local.printable_only
  description = "String with only printable ASCII characters (no tabs, newlines, etc.)"
}

# Example 5: Chaining with Terraform built-ins
locals {
  raw_input = "Café 世界! How are YOU? 🎉"
  
  # Latinize, then use regex to keep only alphanumeric and hyphens, then lowercase
  chained_clean_id = lower(
    replace(
      provider::curious::latinize(local.raw_input),
      "/[^A-Za-z0-9-]/",
      ""
    )
  )
}

output "chained_result" {
  value       = local.chained_clean_id
  description = "Latinized, normalized, and lowercased"
}

# Example 6: Case conversions
locals {
  input_text = "Hello World! Café-Example_123"
  
  flat_case   = provider::curious::flat(local.input_text)
  kebab_case  = provider::curious::kebab(local.input_text)
  camel_case  = provider::curious::camel(local.input_text)
  pascal_case = provider::curious::pascal(local.input_text)
  snake_case  = provider::curious::snake(local.input_text)
  upper_case  = provider::curious::upper(local.input_text)
  train_case  = provider::curious::train(local.input_text)
  ada_case    = provider::curious::ada(local.input_text)
}

output "case_conversions" {
  value = {
    flat   = local.flat_case
    kebab  = local.kebab_case
    camel  = local.camel_case
    pascal = local.pascal_case
    snake  = local.snake_case
    upper  = local.upper_case
    train  = local.train_case
    ada    = local.ada_case
  }
  description = "All case conversion formats"
}

# Example 7: Flavor cases
locals {
  flavor_input = "Sponge Bob! Café åäöñ"
  elite_case   = provider::curious::elite(local.flavor_input)
  sponge_case  = provider::curious::sponge(local.flavor_input)
}

output "flavor_cases" {
  value = {
    elite  = local.elite_case
    sponge = local.sponge_case
  }
  description = "Elite and sponge case examples (elite: SPoNGe BoB! CaFé; sponge: sPoNgE bOb! cAfÉ ÅäÖñ)"
}

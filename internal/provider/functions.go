package provider

import (
	"context"
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// latinize removes diacritical marks from a string
func latinize(input string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, input)
	return result, err
}

// splitWords splits a latinized string into words by non-alphanumeric characters
func splitWords(s string) []string {
	var words []string
	var word strings.Builder
	
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			word.WriteRune(r)
		} else if word.Len() > 0 {
			words = append(words, word.String())
			word.Reset()
		}
	}
	if word.Len() > 0 {
		words = append(words, word.String())
	}
	
	return words
}

func hasDiacritic(r rune) bool {
	if !unicode.IsLetter(r) {
		return false
	}

	decomposed := norm.NFD.String(string(r))
	for _, dr := range decomposed {
		if unicode.Is(unicode.Mn, dr) {
			return true
		}
	}

	return false
}

func isVowel(r rune) bool {
	switch unicode.ToLower(r) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return hasDiacritic(r)
	}
}

// AsciiFunction removes all non-ASCII characters from a string
var _ function.Function = &AsciiFunction{}

type AsciiFunction struct{}

func NewAsciiFunction() function.Function {
	return &AsciiFunction{}
}

func (f *AsciiFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ascii"
}

func (f *AsciiFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Remove non-ASCII characters",
		Description: "Removes diacritics first, then removes all non-ASCII characters from the input string, keeping only characters with ASCII values 0-127.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to process",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *AsciiFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	// First, latinize the input to remove diacritics
	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	// Then remove non-ASCII characters
	var result strings.Builder
	for _, r := range latinized {
		if r <= unicode.MaxASCII {
			result.WriteRune(r)
		}
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result.String()))
}

// AsciiPrintableFunction removes all non-printable ASCII characters from a string
var _ function.Function = &AsciiPrintableFunction{}

type AsciiPrintableFunction struct{}

func NewAsciiPrintableFunction() function.Function {
	return &AsciiPrintableFunction{}
}

func (f *AsciiPrintableFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ascii_printable"
}

func (f *AsciiPrintableFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Remove non-printable ASCII characters",
		Description: "Removes diacritics first, then removes all characters except printable ASCII (32-126), which excludes control characters like tabs and newlines.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to process",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *AsciiPrintableFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	// First, latinize the input to remove diacritics
	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	// Then keep only printable ASCII (32-126)
	var result strings.Builder
	for _, r := range latinized {
		if r >= 32 && r <= 126 {
			result.WriteRune(r)
		}
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result.String()))
}

// LatinizeFunction removes diacritics from a string, converting accented characters to their base Latin equivalents
var _ function.Function = &LatinizeFunction{}

type LatinizeFunction struct{}

func NewLatinizeFunction() function.Function {
	return &LatinizeFunction{}
}

func (f *LatinizeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "latinize"
}

func (f *LatinizeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Remove diacritics (latinize)",
		Description: "Removes diacritical marks (accents) from characters, converting them to their base Latin equivalents. For example: 'räksmörgås' becomes 'raksmorgas'.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to latinize",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *LatinizeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	result, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}

// FlatFunction converts to flatcase (all lowercase, no separators)
var _ function.Function = &FlatFunction{}

type FlatFunction struct{}

func NewFlatFunction() function.Function {
	return &FlatFunction{}
}

func (f *FlatFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "flat"
}

func (f *FlatFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to flatcase",
		Description: "Converts to flatcase: all lowercase with no separators. Latinizes first, then splits on non-alphanumeric characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *FlatFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	words := splitWords(latinized)
	result := strings.ToLower(strings.Join(words, ""))
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}

// KebabFunction converts to kebab-case (lowercase with hyphens)
var _ function.Function = &KebabFunction{}

type KebabFunction struct{}

func NewKebabFunction() function.Function {
	return &KebabFunction{}
}

func (f *KebabFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "kebab"
}

func (f *KebabFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to kebab-case",
		Description: "Converts to kebab-case: lowercase words separated by hyphens. Latinizes first, then splits on non-alphanumeric characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *KebabFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	words := splitWords(latinized)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	result := strings.Join(words, "-")
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}

// CamelFunction converts to camelCase
var _ function.Function = &CamelFunction{}

type CamelFunction struct{}

func NewCamelFunction() function.Function {
	return &CamelFunction{}
}

func (f *CamelFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "camel"
}

func (f *CamelFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to camelCase",
		Description: "Converts to camelCase: first word lowercase, subsequent words capitalized, no separators. Latinizes first, then splits on non-alphanumeric characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *CamelFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	words := splitWords(latinized)
	if len(words) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, ""))
		return
	}

	var result strings.Builder
	result.WriteString(strings.ToLower(words[0]))
	for i := 1; i < len(words); i++ {
		result.WriteString(strings.Title(strings.ToLower(words[i])))
	}
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result.String()))
}

// PascalFunction converts to PascalCase
var _ function.Function = &PascalFunction{}

type PascalFunction struct{}

func NewPascalFunction() function.Function {
	return &PascalFunction{}
}

func (f *PascalFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "pascal"
}

func (f *PascalFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to PascalCase",
		Description: "Converts to PascalCase: all words capitalized, no separators. Latinizes first, then splits on non-alphanumeric characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *PascalFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	words := splitWords(latinized)
	var result strings.Builder
	for _, word := range words {
		result.WriteString(strings.Title(strings.ToLower(word)))
	}
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result.String()))
}

// SnakeFunction converts to snake_case
var _ function.Function = &SnakeFunction{}

type SnakeFunction struct{}

func NewSnakeFunction() function.Function {
	return &SnakeFunction{}
}

func (f *SnakeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "snake"
}

func (f *SnakeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to snake_case",
		Description: "Converts to snake_case: lowercase words separated by underscores. Latinizes first, then splits on non-alphanumeric characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *SnakeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	words := splitWords(latinized)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	result := strings.Join(words, "_")
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}

// UpperFunction converts to UPPER_CASE
var _ function.Function = &UpperFunction{}

type UpperFunction struct{}

func NewUpperFunction() function.Function {
	return &UpperFunction{}
}

func (f *UpperFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "upper"
}

func (f *UpperFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to UPPER_CASE",
		Description: "Converts to UPPER_CASE: uppercase words separated by underscores. Latinizes first, then splits on non-alphanumeric characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *UpperFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	words := splitWords(latinized)
	for i := range words {
		words[i] = strings.ToUpper(words[i])
	}
	result := strings.Join(words, "_")
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}

// TrainFunction converts to TRAIN-CASE
var _ function.Function = &TrainFunction{}

type TrainFunction struct{}

func NewTrainFunction() function.Function {
	return &TrainFunction{}
}

func (f *TrainFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "train"
}

func (f *TrainFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to TRAIN-CASE",
		Description: "Converts to TRAIN-CASE: uppercase words separated by hyphens. Latinizes first, then splits on non-alphanumeric characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *TrainFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	words := splitWords(latinized)
	for i := range words {
		words[i] = strings.ToUpper(words[i])
	}
	result := strings.Join(words, "-")
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}

// AdaFunction converts to Ada_Case
var _ function.Function = &AdaFunction{}

type AdaFunction struct{}

func NewAdaFunction() function.Function {
	return &AdaFunction{}
}

func (f *AdaFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ada"
}

func (f *AdaFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to Ada_Case",
		Description: "Converts to Ada_Case: capitalized words separated by underscores. Latinizes first, then splits on non-alphanumeric characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *AdaFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	latinized, err := latinize(input)
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	words := splitWords(latinized)
	for i := range words {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}
	result := strings.Join(words, "_")
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}

// EliteFunction converts to elite case (uppercase consonants, lowercase vowels)
var _ function.Function = &EliteFunction{}

type EliteFunction struct{}

func NewEliteFunction() function.Function {
	return &EliteFunction{}
}

func (f *EliteFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "elite"
}

func (f *EliteFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to elite case",
		Description: "Uppercases consonants and lowercases vowels, leaving non-letter characters unchanged. Treats letters with diacritics as vowels.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *EliteFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	var result strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) {
			if isVowel(r) {
				result.WriteRune(unicode.ToLower(r))
			} else {
				result.WriteRune(unicode.ToUpper(r))
			}
		} else {
			result.WriteRune(r)
		}
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result.String()))
}

// SpongeFunction converts to sponge case (alternate lowercase/uppercase on letters)
var _ function.Function = &SpongeFunction{}

type SpongeFunction struct{}

func NewSpongeFunction() function.Function {
	return &SpongeFunction{}
}

func (f *SpongeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "sponge"
}

func (f *SpongeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Convert to sponge case",
		Description: "Alternates lowercase and uppercase letters, starting with lowercase for each word. Non-letter characters are unchanged and reset the alternation.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *SpongeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	var result strings.Builder
	useLower := true
	for _, r := range input {
		if unicode.IsLetter(r) {
			if useLower {
				result.WriteRune(unicode.ToLower(r))
			} else {
				result.WriteRune(unicode.ToUpper(r))
			}
			useLower = !useLower
		} else {
			result.WriteRune(r)
			useLower = true
		}
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result.String()))
}

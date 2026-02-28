package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &curiousProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &curiousProvider{
			version: version,
		}
	}
}

// curiousProvider is the provider implementation.
type curiousProvider struct {
	version string
}

// Metadata returns the provider type name.
func (p *curiousProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "curious"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *curiousProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provider for string normalization functions. This provider does not manage any resources.",
	}
}

// Configure prepares the provider for use.
func (p *curiousProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// No configuration needed for this provider
}

// DataSources defines the data sources implemented in the provider.
func (p *curiousProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Resources defines the resources implemented in the provider.
func (p *curiousProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

// Functions defines the functions implemented in the provider.
func (p *curiousProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewAsciiFunction,
		NewAsciiPrintableFunction,
		NewLatinizeFunction,
		NewFlatFunction,
		NewKebabFunction,
		NewCamelFunction,
		NewPascalFunction,
		NewSnakeFunction,
		NewUpperFunction,
		NewTrainFunction,
		NewAdaFunction,
		NewEliteFunction,
		NewSpongeFunction,
	}
}

package curl

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider = &curlProvider{}
)

type curlProvider struct{}

func New() provider.Provider {
	return &curlProvider{}
}

// Metadata returns the provider type name.
func (p *curlProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "curl"
}

// Schema defines the provider-level schema for configuration data.
func (p *curlProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Make curl requests via Terraform.",
		Attributes: map[string]schema.Attribute{
			"token": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares a curl client for data sources and resources.
func (p *curlProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config providerConfiguration
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	token := os.Getenv("CURL_OAUTH2_TOKEN")
	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	opts := HttpClientOptions{token: token}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := NewClient(ctx, opts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create http Client",
			"An unexpected error occurred when creating the http client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"HTTP Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *curlProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCurlRequestDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *curlProvider) Resources(context.Context) []func() resource.Resource {
	return nil
}

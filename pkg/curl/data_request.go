package curl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &curlRequestDataSource{}
	_ datasource.DataSourceWithConfigure = &curlRequestDataSource{}
)

// NewCurlRequestDataSource instantiates a new curl request data source.
func NewCurlRequestDataSource() datasource.DataSource {
	return &curlRequestDataSource{}
}

type curlRequestDataSource struct {
	client *http.Client
}

// Configure implements datasource.DataSourceWithConfigure.
func (d *curlRequestDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*HttpClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	d.client = client.httpClient
}

// Metadata returns the resoure type name.
func (*curlRequestDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_request"
}

type requestDataSourceModel struct {
	ID                 types.Int64  `tfsdk:"id"`
	URI                types.String `tfsdk:"uri"`
	HTTPMethod         types.String `tfsdk:"http_method"`
	Data               types.String `tfsdk:"data"`
	Headers            types.String `tfsdk:"headers"`
	ResponseStatusCode types.Int64  `tfsdk:"response_status_code"`
	ResponseBody       types.String `tfsdk:"response_body"`
}

// Schema implements datasource.DataSource.
func (*curlRequestDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provision a Dex oauth2 client.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"uri": schema.StringAttribute{
				Description: "URI of resource you'd like to retrieve via HTTP(s).",
				Required:    true,
			},
			"http_method": schema.StringAttribute{
				Description: "HTTP method like GET, POST, PUT, DELETE, PATCH.",
				Required:    true,
			},
			"data": schema.StringAttribute{
				Description: "The data sent in the request.",
				Optional:    true,
			},
			"headers": schema.StringAttribute{
				Description: "Headers sent in the request.",
				Optional:    true,
			},
			"response_status_code": schema.Int64Attribute{
				Description: "HTTP Statuscode returned from the HTTP request.",
				Computed:    true,
			},
			"response_body": schema.StringAttribute{
				Description: "Response body returned from the HTTP request.",
				Computed:    true,
			},
		},
	}
}

// Read implements datasource.DataSource.
func (d *curlRequestDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state requestDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data := state.Data.ValueString()

	var body io.Reader
	if strings.Trim(data, " \t\r\n") != "" {
		body = strings.NewReader(data)
	} else {
		body = bytes.NewReader([]byte{})
	}

	httpRequest, err := http.NewRequest(state.HTTPMethod.ValueString(), state.URI.ValueString(), body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create http.Request",
			err.Error(),
		)
		return
	}

	headers := state.Headers.ValueString()

	if strings.Trim(headers, " \t\r\n") != "" {
		parsedHeaders, err := parseJSON(headers)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to set http headers",
				err.Error(),
			)
			return
		}

		for header, value := range parsedHeaders {
			httpRequest.Header.Set(header, value.(string))
		}
	}

	r, err := d.client.Do(httpRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to make http.Request",
			err.Error(),
		)
		return
	}
	defer r.Body.Close()

	if r.StatusCode >= http.StatusBadRequest {
		resp.Diagnostics.AddError(
			"Request failed",
			fmt.Sprintf("http_status: %s", r.Status),
		)
		return
	}

	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed reading response body",
			err.Error(),
		)
		return
	}

	state.ID = types.Int64Value(int64(time.Now().Unix()))
	state.ResponseStatusCode = types.Int64Value(int64(r.StatusCode))
	state.ResponseBody = types.StringValue(string(responseData))

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func parseJSON(input string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

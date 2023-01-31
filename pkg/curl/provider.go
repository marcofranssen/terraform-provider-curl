package curl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func NewProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CURL_OAUTH2_TOKEN", nil),
			},
			"disabletls": {
				Type:      schema.TypeBool,
				Optional:  true,
				Sensitive: false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"curl_request": dataSource(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	disabletls := d.Get("disabletls").(bool)

	var diags diag.Diagnostics
	opts := HttpClientOptions{
		token:      token,
		disabletls: disabletls,
	}

	if opts.token != "" {
		tflog.MaskAllFieldValuesStrings(ctx, token)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c, err := NewClient(ctx, opts)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}

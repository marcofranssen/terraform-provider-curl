package curl

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCurlRequest,
		Schema: map[string]*schema.Schema{
			"uri": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URI of resource you'd like to retrieve via HTTP(s).",
			},
			"http_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "HTTP method like GET, POST, PUT, DELETE, PATCH.",
			},
			"data": {
				Type:        schema.TypeString,
				Optional:    true,
				Required:    false,
				Computed:    false,
				Description: "The data sent in the request.",
			},
			"response_status_code": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "HTTP Statuscode returned from the HTTP request.",
			},
			"response_body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Response body returned from the HTTP request.",
			},
		},
	}
}

func dataSourceCurlRequest(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	myClient := meta.(*HttpClient)

	var diags diag.Diagnostics

	uri := d.Get("uri").(string)
	httpMethod := d.Get("http_method").(string)
	data := d.Get("data").(string)

	var body io.Reader
	if strings.Trim(data, " \t\r\n") != "" {
		body = strings.NewReader(data)
	} else {
		body = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(httpMethod, uri, body)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := myClient.httpClient.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	if r.StatusCode >= http.StatusBadRequest {
		diag.Errorf("Request failed with http status %d, %s")
	}

	if err := d.Set("response_status_code", r.StatusCode); err != nil {
		return diag.FromErr(err)
	}

	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("response_body", string(responseData)); err != nil {
		return diag.FromErr(err)
	}

	// force that it always sets for the newest json object by changing the id of the object
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

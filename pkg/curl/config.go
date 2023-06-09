package curl

import "github.com/hashicorp/terraform-plugin-framework/types"

type providerConfiguration struct {
	Token types.String `tfsdk:"token"`
}

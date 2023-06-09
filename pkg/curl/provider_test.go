package curl_test

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/marcofranssen/terraform-provider-curl/pkg/curl"
)

func GetProviderConfig() string {
	return `
provider "curl" {
	token = "adf12346sdf656d5f66df"
}`
}

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"curl": providerserver.NewProtocol6WithError(curl.New()),
	}
)

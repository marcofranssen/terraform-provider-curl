package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/marcofranssen/terraform-provider-curl/pkg/curl"
)

// Provider documentation generation.
//
//go:generate tfplugindocs generate --provider-name curl
func main() {
	providerserver.Serve(context.Background(), curl.New, providerserver.ServeOpts{
		Address: "hashicorp.com/marcofranssen/curl",
	})
}

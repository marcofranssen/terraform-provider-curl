package curl_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	testResourceName = "data.curl_request.google"
)

func TestCurlRequestDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: GetProviderConfig() + `
data "curl_request" "google" {
	uri         = "https://google.com"
	http_method = "GET"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testResourceName, "uri", "https://google.com"),
					resource.TestCheckResourceAttr(testResourceName, "http_method", "GET"),
					resource.TestCheckResourceAttrWith(testResourceName, "id", func(value string) error {
						id, err := strconv.ParseInt(value, 10, 64)
						if err != nil {
							return err
						}

						diff := time.Since(time.Unix(id, 0)).Seconds()
						if diff > 1.5 {
							return fmt.Errorf("expected ID to be within 1.5 seconds of current unix timestamp, was %f seconds difference", diff)
						}

						return nil
					}),
					resource.TestCheckResourceAttr(testResourceName, "response_status_code", "200"),
					resource.TestCheckResourceAttrWith(testResourceName, "response_body", func(value string) error {
						if !strings.Contains(value, "html") && !strings.Contains(value, "<title>Google</title>") {
							return errors.New("Expected to get html page back")
						}

						return nil
					}),
				),
			},
		},
	})
}

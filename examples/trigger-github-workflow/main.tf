locals {
  ref      = "deploy-eks-workflow"
  workflow = "account-customizations.yaml"
}

// Ensure the curl_request happens every single apply
resource "random_id" "test" {
  keepers = {
    first = "${timestamp()}"
  }
  byte_length = 8
}

data "curl_request" "dispatch_github_account_customizations_workflow" {
  uri         = "https://api.github.com/repos/philips-internal/fiesta-aft-account-request/actions/workflows/${local.workflow}/dispatches"
  http_method = "POST"
  data = jsonencode({
    ref : local.ref,
    inputs : {
      account_id : "315773391160",
      aws_region : "us-east-1",
      enable_eks : "false",
    }
  })

  depends_on = [
    random_id.test
  ]
}

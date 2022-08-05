# Trigger github workflow Example

This example shows how to dispatch a GitHub workflow. This could be in particular useful when chaining various terraform environments living in different repositories to apply them in a particular order.

## Run the example

A trick used to always trigger the curl request to the GitHub workflow is done with the random resource.

```shell
terraform init
terraform apply # invokes the workflow
terraform apply # invokes the workflow
terraform apply # invokes the workflow
```

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.1.9 |
| <a name="requirement_curl"></a> [curl](#requirement\_curl) | ~> 0.2.0 |
| <a name="requirement_random"></a> [random](#requirement\_random) | 3.3.2 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_curl"></a> [curl](#provider\_curl) | 0.1.0 |
| <a name="provider_random"></a> [random](#provider\_random) | 3.3.2 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [random_id.test](https://registry.terraform.io/providers/hashicorp/random/3.3.2/docs/resources/id) | resource |
| [curl_request.dispatch_github_account_customizations_workflow](https://registry.terraform.io/providers/marcofranssen/curl/latest/docs/data-sources/request) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_github_token"></a> [github\_token](#input\_github\_token) | The github token | `string` | n/a | yes |
| <a name="input_owner"></a> [owner](#input\_owner) | The github repository owner | `string` | n/a | yes |
| <a name="input_ref"></a> [ref](#input\_ref) | The git reference to target the workflow dispatch at | `string` | `"main"` | no |
| <a name="input_repo"></a> [repo](#input\_repo) | The github repository | `string` | n/a | yes |
| <a name="input_workflow"></a> [workflow](#input\_workflow) | The name of the workflow (e.g. deploy.yaml) | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_workflow_response"></a> [workflow\_response](#output\_workflow\_response) | n/a |
<!-- END_TF_DOCS -->
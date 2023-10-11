# ifconfig.co Example

This example shows how to get your public IP information. This could be usefull if you for example need the outbound IP address from the location where you are running Terraform apply.

## Run the example

A trick used to always trigger the curl request to ifconfig.co is done with the random resource.

```shell
terraform init
terraform apply # Get your public ip information from ifconfig.co
terraform apply # Get your public ip information from ifconfig.co
terraform apply # Get your public ip information from ifconfig.co
```

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.1.9 |
| <a name="requirement_curl"></a> [curl](#requirement\_curl) | ~> 0.4.1 |
| <a name="requirement_random"></a> [random](#requirement\_random) | 3.3.2 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_curl"></a> [curl](#provider\_curl) | 0.2.1 |
| <a name="provider_random"></a> [random](#provider\_random) | 3.3.2 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [random_id.test](https://registry.terraform.io/providers/hashicorp/random/3.3.2/docs/resources/id) | resource |
| [curl_request.ifconfig](https://registry.terraform.io/providers/marcofranssen/curl/latest/docs/data-sources/request) | data source |

## Inputs

No inputs.

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_ifconfig_response"></a> [ifconfig\_response](#output\_ifconfig\_response) | n/a |
<!-- END_TF_DOCS -->

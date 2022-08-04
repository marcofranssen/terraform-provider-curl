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

# ifconfig.co Example

This example shows how to dispatch a GitHub workflow. This could be in particular usefull when chaining various terraform environments living in different repositories to apply them in a particular order.

## Run the example

A trick used to always trigger the curl request to the GitHub workflow is done with the random resource.

```shell
terraform init
terraform apply # invokes the workflow
terraform apply # invokes the workflow
terraform apply # invokes the workflow
```

output "workflow_response" {
  value = {
    status_code = data.curl_request.dispatch_github_account_customizations_workflow.response_status_code
    body        = data.curl_request.dispatch_github_account_customizations_workflow.response_body
  }
}

output "ifconfig_response" {
  value = {
    status_code = data.curl_request.ifconfig.response_status_code
    body        = jsondecode(data.curl_request.ifconfig.response_body)
  }
}

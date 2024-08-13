data "curl_request" "ipify" {
  uri         = "https://api.ipify.org?format=json"
  http_method = "GET"
  headers = jsonencode({
    Content-Type = "application/json"
  })
}
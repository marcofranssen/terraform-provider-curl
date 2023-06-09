data "curl_request" "google" {
    uri         = "https://google.com"
    http_method = "GET"
}

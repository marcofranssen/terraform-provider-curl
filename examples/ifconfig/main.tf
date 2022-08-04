// Ensure the curl_request happens every single apply
resource "random_id" "test" {
  keepers = {
    first = "${timestamp()}"
  }
  byte_length = 8
}

data "curl_request" "ifconfig" {
  uri         = "https://ifconfig.co/json"
  http_method = "GET"

  depends_on = [
    random_id.test
  ]
}

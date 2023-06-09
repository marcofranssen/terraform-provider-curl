// unauthenticated provider
provider "curl" {}

// provider with a oauth2 token to be used for requests
provider "curl" {
    token = ""
}

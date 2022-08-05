terraform {
  required_version = ">= 1.1.9"

  required_providers {
    curl = {
      source  = "marcofranssen/curl"
      version = "~> 0.2.1"
    }
  }
}

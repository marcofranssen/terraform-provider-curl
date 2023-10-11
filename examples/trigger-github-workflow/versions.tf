terraform {
  required_version = ">= 1.1.9"

  required_providers {
    curl = {
      source  = "marcofranssen/curl"
      version = "~> 0.4.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.3.2"
    }
  }
}

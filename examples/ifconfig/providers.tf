provider "curl" {
  token = replace(file("/Volumes/Keybase/team/swatehv/fiesta/terraform-pat.txt"), "\n", "")
}

provider "random" {
  # Configuration options
}

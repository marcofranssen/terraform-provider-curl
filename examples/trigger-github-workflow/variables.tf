variable "github_token" {
  description = "The github token"
  type        = string
  sensitive   = true
}

variable "owner" {
  description = "The github repository owner"
  type        = string
}

variable "repo" {
  description = "The github repository"
  type        = string
}

variable "workflow" {
  description = "The name of the workflow (e.g. deploy.yaml)"
  type        = string
}

variable "ref" {
  description = "The git reference to target the workflow dispatch at"
  type        = string
  default     = "main"
}


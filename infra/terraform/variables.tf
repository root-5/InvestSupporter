variable "project_id" {
  description = "GCP Project ID"
  type        = string
}

variable "region" {
  description = "GCP Region"
  type        = string
  default     = "asia-northeast1"
}

variable "zone" {
  description = "GCP Zone"
  type        = string
  default     = "asia-northeast1-a"
}

variable "machine_type" {
  description = "GCE Machine Type"
  type        = string
  default     = "e2-medium"
}

variable "instance_name" {
  description = "GCE Instance Name"
  type        = string
  default     = "invest-supporter-app"
}

variable "allowed_ssh_ips" {
  description = "List of IP addresses allowed to SSH"
  type        = list(string)
  default     = ["0.0.0.0/0"] # セキュリティのため、運用時は特定のIPに制限することを推奨
}

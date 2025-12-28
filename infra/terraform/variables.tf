variable "project_id" {
  description = "プロジェクト ID"
  type        = string
}

variable "region" {
  description = "リージョン"
  type        = string
  default     = "asia-northeast1"
}

variable "zone" {
  description = "ゾーン"
  type        = string
  default     = "asia-northeast1-a"
}

variable "machine_type" {
  description = CE タイプ"
  type        = string
  default     = "e2-micro"
}

variable "instance_name" {
  description = "インスタンス名"
  type        = string
  default     = "invest-supporter-app"
}

variable "allowed_ssh_ips" {
  description = "SSH 接続を許可する IP リスト"
  type        = list(string)
  default     = ["0.0.0.0/0"] # セキュリティのため、運用時は特定のIPに制限することを推奨
}

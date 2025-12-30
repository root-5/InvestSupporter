# variables 変数の定義
# 変数の宣言・型・説明・デフォルト値を定義、実際の値は terraform.tfvars に記載し git 管理から除外する
# terraform.tfvars 以外にも .auto.tfvars やコマンドライン引数など様々な方法で値を渡せる
# 定義した値を参照したいときは var.<variable_name> でアクセス可能、型情報そのままに参照できる
# 文字列参照したいときは "${var.<variable_name>}_test" のように記述する
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
  description = "GCE タイプ"
  type        = string
  default     = "e2-micro" # ARM インスタンスは規模が大きいものしか提供されていなかったので e2-micro を選択
}

variable "oslogin_members" {
  description = "OS Login を付与するメンバー一覧 (例: user:you@example.com)"
  type        = list(string)
  default     = []
}

variable "osadminlogin_members" {
  description = "OS Admin Login を付与するメンバー一覧 (例: user:you@example.com)"
  type        = list(string)
  default     = []
}

variable "workload_identity_pool_id" {
  description = "Workload Identity Pool の ID"
  type        = string
  default     = "github-actions-pool"
}

variable "workload_identity_provider_id" {
  description = "Workload Identity Provider の ID"
  type        = string
  default     = "github-actions-provider"
}

variable "github_repository" {
  description = "GitHub Actions で許可するリポジトリ (owner/name)"
  type        = string
  default     = "root-5/InvestSupporter"
}

variable "github_allowed_refs" {
  description = "GitHub Actions で許可する ref のリスト"
  type        = list(string)
  default     = ["refs/heads/main"]
}

variable "actions_service_account_id" {
  description = "GitHub Actions 用サービスアカウントのアカウント ID (メールではなく ID 部分)"
  type        = string
  default     = "github-actions-service-account"
}

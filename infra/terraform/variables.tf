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

variable "allowed_ssh_ips" {
  description = "SSH 接続を許可する IP リスト"
  type        = list(string)
  default     = ["0.0.0.0/0"] # セキュリティのため、運用時は特定のIPに制限
}

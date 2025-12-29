# Terraform はエントリーポイントが言語的には存在しない (.tf は設定ファイル群として扱われる) 
# ただし、慣例的に main.tf がメインファイルとして扱われることが多い
# main.tf だけで完結させてもよいが、役割ごとにファイルをモジュールに分割することが多い
terraform {
  # Terraform 本体のバージョン指定、.terraform-version も同様に設定
  required_version = ">= 1.14.3"

  # プロバイダー指定、「拡張機能」に近い概念、各クラウドの API を操作するための宣言
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 7.14.1"
    }
  }

  # 状態管理 (tfstate) の保存先設定
  # バケット名はプロジェクト外を含めてユニークである必要あり
  # インフラ作成 (terraform apply) 時点では存在している必要があるため、事前に手動で作成しておくこと
  # パブリックアクセス不可かつバージョニング有効化を強く推奨
  backend "gcs" {
    bucket  = "invest-supporter"
    prefix  = "terraform/state"
  }
}

# プロバイダーの設定情報
provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

terraform {
  # Terraform のバージョン指定、.terraform-version も同様に設定
  required_version = ">= 1.14.3"

  # プロバイダー指定、「拡張機能」に近い概念で各クラウドの API を操作するための宣言
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }

  # GCS Backend configuration
  # バケット名はユニークである必要があります。作成したバケット名に書き換えてください。
  # backend "gcs" {
  #   bucket  = "invest-supporter-tfstate"
  #   prefix  = "terraform/state"
  # }
}

# プロバイダーの設定情報
provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

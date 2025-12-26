terraform {
  required_version = ">= 1.9.0"
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

provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

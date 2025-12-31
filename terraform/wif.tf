# Workload Identity Federation (WIF) を使用して GitHub Actions から GCP へデプロイ
# https://docs.cloud.google.com/iam/docs/workload-identity-federation-with-deployment-pipelines?hl=ja

# Workload Identity プール
resource "google_iam_workload_identity_pool" "github" {
  project                   = var.project_id
  workload_identity_pool_id = var.workload_identity_pool_id
  display_name              = "GitHub Actions Pool"
  description               = "GitHub Actions OIDC"
}

# Workload Identity プロバイダー (GitHub OIDC)
resource "google_iam_workload_identity_pool_provider" "github" {
  project                            = var.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.github.workload_identity_pool_id
  workload_identity_pool_provider_id = var.workload_identity_provider_id
  display_name                       = "GitHub Actions Provider"

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }

  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
    "attribute.ref"        = "assertion.ref"
    "attribute.actor"      = "assertion.actor"
  }

  attribute_condition = format(
    "attribute.repository==\"%s\" && attribute.ref in %s",
    var.github_repository,
    jsonencode(var.github_allowed_refs)
  )
}

# 現在のプロジェクト情報を取得
data "google_project" "current" {}

# GitHub Actions 用サービスアカウント
resource "google_service_account" "github_actions" {
  account_id   = var.actions_service_account_id
  display_name = "GitHub Actions deploy"
  project      = var.project_id
}

# インスタンスのサービスアカウントを実行できる権限を付与
resource "google_service_account_iam_member" "github_actions_act_as_instance_sa" {
  service_account_id = format(
    "projects/%s/serviceAccounts/%s-compute@developer.gserviceaccount.com",
    var.project_id,
    data.google_project.current.number,
  )
  role   = "roles/iam.serviceAccountUser"
  member = google_service_account.github_actions.member
}

# サービスアカウントへの Workload Identity ユーザー権限付与
resource "google_service_account_iam_member" "github_actions_iam_workload_identity_user" {
  service_account_id = google_service_account.github_actions.name
  role               = "roles/iam.workloadIdentityUser"
  member = format(
    "principalSet://iam.googleapis.com/projects/%s/locations/global/workloadIdentityPools/%s/attribute.repository/%s",
    data.google_project.current.number,
    google_iam_workload_identity_pool.github.workload_identity_pool_id,
    var.github_repository
  )
}

# デプロイに必要なロール（IAP トンネル + OS Login sudo）
resource "google_project_iam_member" "github_actions_iap" {
  project = var.project_id
  role    = "roles/iap.tunnelResourceAccessor"
  member  = google_service_account.github_actions.member
}

resource "google_project_iam_member" "github_actions_os_admin" {
  project = var.project_id
  role    = "roles/compute.osAdminLogin"
  member  = google_service_account.github_actions.member
}

resource "google_project_iam_member" "github_actions_compute_viewer" {
  project = var.project_id
  role    = "roles/compute.viewer"
  member  = google_service_account.github_actions.member
}
